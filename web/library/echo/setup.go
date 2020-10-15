package setup

import (
	"github.com/Seven4X/link/web/library/echo/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewEcho() (e *echo.Echo) {
	// Echo instance
	e = echo.New()
	e.Validator = validator.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//sentinel参考：https://github.com/alibaba/sentinel-golang/tree/master/adapter/echo
	//https://github.com/alibaba/sentinel-golang/blob/master/example/datasource/nacos/datasource_nacos_example.go

	return e
}
