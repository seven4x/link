package server

import (
	"github.com/Seven4X/link/web/app/vote/server/request"
	"github.com/Seven4X/link/web/app/vote/service"
	"github.com/Seven4X/link/web/library/api"
	"github.com/Seven4X/link/web/library/api/messages"
	"github.com/Seven4X/link/web/library/consts"
	"github.com/Seven4X/link/web/library/echo/mymw"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	svr = service.NewService()
)

func Router(e *echo.Echo) {
	e.POST("/vote", vote, mymw.JWT())
}

func vote(e echo.Context) error {

	//todo 重复代码 封装
	req := new(request.VoteRequest)
	if err := e.Bind(req); err != nil {
		e.JSON(http.StatusOK, api.Fail(err.Error()))
		return nil
	}
	if err := e.Validate(req); err != nil {
		e.JSON(http.StatusOK, api.Fail(err.Error()))
		return nil
	}
	u := e.Get(consts.User)
	if u == nil {
		e.JSON(http.StatusOK, api.FailMsgId(messages.GlobalActionMustLogin))
		return nil
	}
	//
	user := u.(*jwt.Token)
	claims := user.Claims.(*mymw.JwtCustomClaims)
	req.CreateBy = claims.Id

	res, err := svr.Vote(req)
	e.JSON(http.StatusOK, api.Response(res, err))

	return nil
}
