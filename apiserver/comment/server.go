package comment

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/seven4x/link/web/messages"

	"github.com/seven4x/link/web/middleware"
	"github.com/seven4x/link/web/util"
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
	g.POST("/:lid/comment", postComment, middleware.JWT())
	g.GET("/:lid/comment", listComment, middleware.Anonymous())
	g.DELETE("/:lid/comment/:mid", deleteComment, middleware.JWT())
}
func postComment(e echo.Context) error {
	req := new(NewCommentRequest)
	if err := e.Bind(req); err != nil {
		e.JSON(http.StatusOK, util.Fail(err.Error()))
		return nil
	}
	if err := e.Validate(req); err != nil {
		e.JSON(http.StatusOK, util.Fail(err.Error()))
		return nil
	}
	u := e.Get(util.User)
	if u == nil {
		e.JSON(http.StatusOK, util.FailMsgId(messages.GlobalActionMustLogin))
		return nil
	}
	user := u.(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	req.CreateBy = claims.Id
	lid := e.Param("lid")
	if id, err := strconv.Atoi(lid); err != nil {
		e.JSON(http.StatusOK, util.FailMsgId(messages.GlobalParamWrong))
	} else {
		req.LinkId = id
	}

	id, err := svr.SaveNewComment(req)
	_ = e.JSON(http.StatusOK, util.Response(id, err))
	return nil
}
func listComment(e echo.Context) error {
	linkId := e.Param("lid")
	pre := e.Param("prev")
	linkIdInt, err := strconv.Atoi(linkId)
	if linkId == "" || err != nil {
		e.JSON(http.StatusOK, util.FailMsgId(messages.GlobalParamWrong))
		return nil
	}
	req := new(ListCommentRequest)
	e.Bind(req)
	req.LinkId = linkIdInt
	req.Size = DefaultSize
	req.UserId = util.GetUserId(e)
	req.Prev, _ = strconv.Atoi(pre)
	res, hasMore, err := svr.ListComment(req)
	if err != nil {
		return err
	}
	e.JSON(http.StatusOK, util.ResponseHasMore(res, hasMore))
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
	e.JSON(http.StatusOK, util.Success(update))
	return nil
}
