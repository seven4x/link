package server

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/seven4x/link/api"
	"github.com/seven4x/link/app"
	"net/http"
)

func RouterVote(e *echo.Echo) {
	e.POST("/api1/vote", Vote, app.JWT())
}

func Vote(e echo.Context) error {

	//todo 重复代码 封装
	req := new(api.VoteRequest)
	if err := e.Bind(req); err != nil {
		e.JSON(http.StatusOK, api.Fail(err.Error()))
		return nil
	}
	req.Type = string(req.TypeCode[0])
	if err := e.Validate(req); err != nil {
		e.JSON(http.StatusOK, api.Fail(err.Error()))
		return nil
	}
	u := e.Get(app.User)
	if u == nil {
		e.JSON(http.StatusOK, api.FailMsgId(api.GlobalActionMustLogin))
		return nil
	}
	//
	user := u.(*jwt.Token)
	claims := user.Claims.(*app.JwtCustomClaims)
	req.CreateBy = claims.Id

	res, err := svr.Vote(req)
	e.JSON(http.StatusOK, api.Response(res, err))

	return nil
}
