package main

import (
	"net/http"

	account "github.com/Seven4X/link/web/app/account/server"
	link "github.com/Seven4X/link/web/app/link/server"
	topic "github.com/Seven4X/link/web/app/topic/server/http"
	setup "github.com/Seven4X/link/web/library/echo"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo/v4"
)

func main() {
	e := setup.NewEcho()
	// Routes
	e.GET("/", hello)

	// 初始化模块
	topic.Router(e)
	account.Router(e)
	link.Router(e)

	// Start server
	e.Server.Addr = ":1323"
	//参考：https://echo.labstack.com/cookbook/graceful-shutdown
	e.Logger.Fatal(gracehttp.Serve(e.Server))

}
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
