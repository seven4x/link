package comment

import (
	"github.com/Seven4X/link/web/app"
	"github.com/Seven4X/link/web/lib/api"
	"github.com/Seven4X/link/web/lib/api/messages"
	"github.com/Seven4X/link/web/lib/consts"
	"github.com/Seven4X/link/web/lib/echo/mymw"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var (
	svr = NewService()
)

const (
	DefaultSize = 10
)

func Router(e *echo.Echo) {
	g := e.Group("/api1/link")
	g.POST("/:lid/comment", postComment, mymw.JWT())
	g.GET("/:lid/comment", listComment)
	g.DELETE("/:lid/comment/:mid", deleteComment, mymw.JWT())
}
func postComment(e echo.Context) error {
	req := new(NewCommentRequest)
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
	user := u.(*jwt.Token)
	claims := user.Claims.(*mymw.JwtCustomClaims)
	req.CreateBy = claims.Id
	lid := e.Param("lid")
	if id, err := strconv.Atoi(lid); err != nil {
		e.JSON(http.StatusOK, api.FailMsgId(messages.GlobalParamWrong))
	} else {
		req.LinkId = id
	}

	id, err := svr.SaveNewComment(req)
	_ = e.JSON(http.StatusOK, api.Response(id, err))
	return nil
}
func listComment(e echo.Context) error {
	linkId := e.Param("lid")
	pre := e.Param("prev")
	linkIdInt, err := strconv.Atoi(linkId)
	if linkId == "" || err != nil {
		e.JSON(http.StatusOK, api.FailMsgId(messages.GlobalParamWrong))
		return nil
	}
	req := new(ListCommentRequest)
	e.Bind(req)
	req.LinkId = linkIdInt
	req.Size = DefaultSize
	req.UserId = app.GetUserId(e)
	req.Prev, _ = strconv.Atoi(pre)
	res, hasMore, err := svr.ListComment(req)
	if err != nil {
		return err
	}
	e.JSON(http.StatusOK, api.ResponseHasMore(res, hasMore))
	return nil
}
func deleteComment(e echo.Context) error {
	lid := e.Param("lid")
	mid := e.Param("mid")
	linkId, err := strconv.Atoi(lid)
	if err != nil {
		return err
	}
	commentId, err := strconv.Atoi(mid)
	if err != nil {
		return err
	}
	update, err := svr.DeleteComment(linkId, commentId)
	if err != nil {
		return err
	}
	e.JSON(http.StatusOK, api.Success(update))
	return nil
}
