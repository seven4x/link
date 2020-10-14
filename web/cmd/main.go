package main

import (
	"net/http"

	topic "github.com/Seven4X/link/web/app/topic/server/http"
	setup "github.com/Seven4X/link/web/library/echo"
	"github.com/labstack/echo/v4"
)

func main() {
	e := setup.NewEcho()
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
