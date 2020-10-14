package server

import (
	"github.com/Seven4X/link/web/app/account/api/request"
	"github.com/Seven4X/link/web/app/account/service"
	"github.com/Seven4X/link/web/library/api"
	"github.com/Seven4X/link/web/library/echo/mymw"
	"github.com/Seven4X/link/web/library/log"
	"github.com/dgrijalva/jwt-go"
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
	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(*mymw.JwtCustomClaims)
	log.Info(claims)
	e.JSON(http.StatusOK, api.Succ([2]interface{}{"鸡要文件", claims}))
	return nil
}
