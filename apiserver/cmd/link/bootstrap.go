package main

import (
	adapter "github.com/alibaba/sentinel-golang/adapter/echo"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/robfig/cron/v3"
	"github.com/seven4x/link/web/comment"
	"github.com/seven4x/link/web/job"
	"github.com/seven4x/link/web/link"
	mymd "github.com/seven4x/link/web/middleware"
	"github.com/seven4x/link/web/topic"
	"github.com/seven4x/link/web/user"
	"github.com/seven4x/link/web/util"
	"github.com/seven4x/link/web/vote"
	"net/http"
	"strings"
	"time"
)

func SetupEcho() (e *echo.Echo) {
	// Echo instance
	e = echo.New()
	e.Validator = mymd.New()
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
		util.Error("initSentinel-error", err.Error())
	}

}

func initRouter(e *echo.Echo) {

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
			util.Error(err.Error())
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
