package setup

import (
	"github.com/Seven4X/link/web/app/comment"
	"github.com/Seven4X/link/web/app/job"
	"github.com/Seven4X/link/web/app/link"
	"github.com/Seven4X/link/web/app/topic"
	"github.com/Seven4X/link/web/app/user"
	"github.com/Seven4X/link/web/app/vote"
	"github.com/Seven4X/link/web/lib/config"
	"github.com/Seven4X/link/web/lib/log"
	"github.com/Seven4X/link/web/lib/setup/validator"
	"github.com/Seven4X/link/web/lib/util"
	adapter "github.com/alibaba/sentinel-golang/adapter/echo"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/robfig/cron/v3"
	"net/http"
	"strings"
	"time"
)

func SetupEcho() (e *echo.Echo) {
	// Echo instance
	e = echo.New()
	e.Validator = validator.New()
	logConfig := middleware.LoggerConfig{}
	// 忽略静态文件日志
	logConfig.Skipper = func(e echo.Context) bool {
		return strings.HasPrefix(e.Path(), "/static") || strings.HasPrefix(e.Path(), "/favicon.ico") || strings.HasPrefix(e.Path(), "/manifest.json")
	}
	e.Use(middleware.LoggerWithConfig(logConfig))
	e.Use(middleware.Recover())

	//sentinel参考：https://github.com/alibaba/sentinel-golang/tree/master/adapter/echo
	//https://github.com/alibaba/sentinel-golang/blob/master/example/datasource/nacos/datasource_nacos_example.go

	initSentinel()
	//全局限流
	e.Use(adapter.SentinelMiddleware())
	//ip限流
	e.Use(
		adapter.SentinelMiddleware(
			// customize resource extractor if required
			// method_path by default
			adapter.WithResourceExtractor(func(ctx echo.Context) string {
				if res, ok := ctx.Get("X-Real-IP").(string); ok {
					return res
				}
				return ""
			}),
			// customize block fallback if required
			// abort with status 429 by default
			adapter.WithBlockFallback(func(ctx echo.Context) error {
				return ctx.JSON(400, map[string]interface{}{
					"err":  "too many requests; the quota used up",
					"code": 10222,
				})
			}),
		),
	)

	e.HTTPErrorHandler = customHTTPErrorHandler

	initRouter(e)
	return e
}
func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	res := map[string]interface{}{"ok": false, "msg": err.Error()}
	if err := c.JSON(code, res); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}

func initSentinel() {
	err := sentinel.InitDefault()
	if err != nil {
		// 初始化 Sentinel 失败
		log.Error("initSentinel-error", err.Error())
	}

}

func initRouter(e *echo.Echo) {
	//站点静态文件build输出目录
	path := config.GetString("site.path")
	e.File("/*", path+"/index.html")
	e.Static("/static", path+"/static")
	e.File("/favicon.ico", path+"/favicon.ico")
	e.File("/manifest.json", path+"/manifest.json")
	//用于证书认证 https://letsencrypt.org/zh-cn/docs/challenge-types/
	e.Static("/.well-known", path+"/wellknown")
	// 初始化模块
	topic.Router(e)
	user.Router(e)
	link.Router(e)
	vote.Router(e)
	comment.Router(e)
}

func StartJob() *cron.Cron {
	local, _ := time.LoadLocation("Local")
	c := cron.New(cron.WithLocation(local))
	c.AddFunc("@midnight", func() {
		err := job.RefreshHotTopic()
		if err != nil {
			log.Error(err.Error())
		}
	})
	c.AddFunc("@hourly", func() {
		util.DumpCuckooFilter()
	})
	//c.AddFunc("@every 10s", func() {
	//	log.Info("活")
	//})
	c.Start()

	return c
}
