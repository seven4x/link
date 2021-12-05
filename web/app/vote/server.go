package vote

import (
	"github.com/Seven4X/link/web/app/messages"
	"github.com/Seven4X/link/web/app/middleware"
	"github.com/Seven4X/link/web/app/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	svr = NewService()
)

func Router(e *echo.Echo) {
	e.POST("/api1/vote", vote, middleware.JWT())
}

func vote(e echo.Context) error {

	//todo 重复代码 封装
	req := new(VoteRequest)
	if err := e.Bind(req); err != nil {
		e.JSON(http.StatusOK, util.Fail(err.Error()))
		return nil
	}
	req.Type = string(req.TypeCode[0])
	if err := e.Validate(req); err != nil {
		e.JSON(http.StatusOK, util.Fail(err.Error()))
		return nil
	}
	u := e.Get(util.User)
	if u == nil {
		e.JSON(http.StatusOK, util.FailMsgId(messages.GlobalActionMustLogin))
		return nil
	}
	//
	user := u.(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	req.CreateBy = claims.Id

	res, err := svr.Vote(req)
	e.JSON(http.StatusOK, util.Response(res, err))

	return nil
}
