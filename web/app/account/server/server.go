package server

import (
	"github.com/Seven4X/link/web/app/account/api/request"
	"github.com/Seven4X/link/web/app/account/service"
	"github.com/Seven4X/link/web/library/api"
	"github.com/Seven4X/link/web/library/echo/mymw"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	svr = service.New()
)

func Router(e *echo.Echo) {
	g := e.Group("/account")
	g.POST("/login", login)
	g.GET("/info", info, mymw.JWT())

}

// todo global error handler
func login(e echo.Context) error {
	req := &request.Login{}
	if err := e.Bind(req); err != nil {
		e.JSON(http.StatusOK, api.Fail(err.Error()))
		return nil
	}

	if b, s := svr.Login(*req); b {
		e.JSON(http.StatusOK, api.Succ(s))
	} else {
		e.JSON(http.StatusOK, api.Fail(s))
	}

	return nil
}

func info(e echo.Context) error {
	e.JSON(http.StatusOK, api.Succ("鸡要文件"))
	return nil
}
