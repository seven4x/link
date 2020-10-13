package main

import (
	topic "github.com/Seven4X/link/web/app/topic/server/http"
	"github.com/Seven4X/link/web/library/echo/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()
	e.Validator = validator.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// 初始化模块
	topic.Router(e)

	// Start server
	//todo gracehttp
	e.Logger.Fatal(e.Start(":1323"))
}
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
