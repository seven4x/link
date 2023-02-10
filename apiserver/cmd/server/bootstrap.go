package main

import (
	adapter "github.com/alibaba/sentinel-golang/adapter/echo"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/seven4x/link/app"
	"github.com/seven4x/link/app/log"
	"github.com/seven4x/link/server"
	"net/http"
	"strings"
)

func BootApp(e *echo.Echo) error {

	e.Validator = app.NewCustomValidator()
	logConfig := middleware.LoggerConfig{}
	// 忽略静态文件日志
	logConfig.Skipper = func(e echo.Context) bool {
		return strings.HasPrefix(e.Path(), "/static") || strings.HasPrefix(e.Path(), "/favicon.ico") || strings.HasPrefix(e.Path(), "/manifest.json")
	}
	e.Use(middleware.LoggerWithConfig(logConfig))
	e.Use(middleware.Recover())

	//sentinel参考：https://github.com/alibaba/sentinel-golang/tree/master/adapter/echo
	//https://github.com/alibaba/sentinel-golang/blob/master/example/datasource/nacos/datasource_nacos_example.go
	initSentinel(e)

	e.HTTPErrorHandler = customHTTPErrorHandler

	newServer := server.NewServer(e)
	newServer.InitRouter()

	err := app.Migration()
	return err
}

func initSentinel(e *echo.Echo) {
	err := sentinel.InitDefault()
	if err != nil {
		// 初始化 Sentinel 失败
		log.Error("initSentinel-error", err.Error())
	}
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
