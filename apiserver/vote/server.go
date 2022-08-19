package vote

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"github.com/seven4x/link/web/messages"
	"github.com/seven4x/link/web/middleware"
	"github.com/seven4x/link/web/util"
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
