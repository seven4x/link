package setup

import (
	"github.com/Seven4X/link/web/lib/config"
	"github.com/Seven4X/link/web/lib/echo/validator"
	"github.com/Seven4X/link/web/lib/log"
	adapter "github.com/alibaba/sentinel-golang/adapter/echo"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/ext/datasource"
	"github.com/alibaba/sentinel-golang/ext/datasource/nacos"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
)

func NewEcho() (e *echo.Echo) {
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
	//从acm加载配置
	//rule配置参考flow.rule
	h := datasource.NewFlowRulesHandler(datasource.FlowRuleJsonArrayParser)
	client := config.GetAcmClient()
	nds, err := nacos.NewNacosDataSource(client, "link-hub-go", "flow", h)
	if err != nil {
		log.Warnf("Fail to create nacos data source client, err: %+v", err)
		return
	}
	err = nds.Initialize()
	if err != nil {
		log.Warnf("Fail to initialize nacos data source client, err: %+v", err)
		return
	}
}
