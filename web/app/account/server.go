package account

import (
	"github.com/Seven4X/link/web/library/api"
	"github.com/Seven4X/link/web/library/echo/mymw"
	"github.com/Seven4X/link/web/library/log"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	svr = NewService()
)

func Router(e *echo.Echo) {
	g := e.Group("/account")
	g.POST("/login", login)
	g.GET("/info", info, mymw.JWT())

}

func login(e echo.Context) error {
	req := &Login{}
	if err := e.Bind(req); err != nil {
		e.JSON(http.StatusOK, api.Fail(err.Error()))
		return nil
	}

	if data, err := svr.Login(*req); err == nil {
		e.JSON(http.StatusOK, api.Success(data))
	} else {
		e.JSON(http.StatusOK, api.Fail(err.Error()))
	}

	return nil
}

func info(e echo.Context) error {
	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(*mymw.JwtCustomClaims)
	log.Info(claims)
	e.JSON(http.StatusOK, api.Success([2]interface{}{"鸡要文件", claims}))
	return nil
}
