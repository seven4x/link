package vote

import (
	"github.com/Seven4X/link/web/lib/api"
	"github.com/Seven4X/link/web/lib/api/messages"
	"github.com/Seven4X/link/web/lib/consts"
	"github.com/Seven4X/link/web/lib/echo/mymw"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	svr = NewService()
)

func Router(e *echo.Echo) {
	e.POST("/api1/vote", vote, mymw.JWT())
}

func vote(e echo.Context) error {

	//todo 重复代码 封装
	req := new(VoteRequest)
	if err := e.Bind(req); err != nil {
		e.JSON(http.StatusOK, api.Fail(err.Error()))
		return nil
	}
	req.Type = string(req.TypeCode[0])
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
