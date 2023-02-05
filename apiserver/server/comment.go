package server

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/seven4x/link/api"
	"github.com/seven4x/link/app"
	"net/http"
	"strconv"
)

func (s *Server) RouterComment() {
	e := s.echo
	g := e.Group("/api1/link")
	g.POST("/:lid/comment", s.PostComment, app.JWT())
	g.GET("/:lid/comment", s.ListComment, app.Anonymous())
	g.DELETE("/:lid/comment/:mid", s.DeleteComment, app.JWT())
}

func (s *Server) PostComment(e echo.Context) error {
	req := new(api.NewCommentRequest)
	if err := e.Bind(req); err != nil {
		e.JSON(http.StatusOK, api.Fail(err.Error()))
		return nil
	}
	if err := e.Validate(req); err != nil {
		e.JSON(http.StatusOK, api.Fail(err.Error()))
		return nil
	}
	u := e.Get(app.User)
	if u == nil {
		e.JSON(http.StatusOK, api.FailMsgId(api.GlobalActionMustLogin))
		return nil
	}
	user := u.(*jwt.Token)
	claims := user.Claims.(*app.JwtCustomClaims)
	req.CreateBy = claims.Id
	lid := e.Param("lid")
	if id, err := strconv.Atoi(lid); err != nil {
		e.JSON(http.StatusOK, api.FailMsgId(api.GlobalParamWrong))
	} else {
		req.LinkId = id
	}

	id, err := s.svr.SaveNewComment(req)
	_ = e.JSON(http.StatusOK, api.Response(id, err))
	return nil
}
func (s *Server) ListComment(e echo.Context) error {
	linkId := e.Param("lid")
	pre := e.Param("prev")
	linkIdInt, err := strconv.Atoi(linkId)
	if linkId == "" || err != nil {
		e.JSON(http.StatusOK, api.FailMsgId(api.GlobalParamWrong))
		return nil
	}
	req := new(api.ListCommentRequest)
	e.Bind(req)
	req.LinkId = linkIdInt
	req.Size = DefaultSize
	req.UserId = app.GetUserId(e)
	req.Prev, _ = strconv.Atoi(pre)
	res, hasMore, err := s.svr.ListComment(req)
	if err != nil {
		return err
	}
	e.JSON(http.StatusOK, api.ResponseHasMore(res, hasMore))
	return nil
}
func (s *Server) DeleteComment(e echo.Context) error {
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
	update, err := s.svr.DeleteComment(linkId, commentId)
	if err != nil {
		return err
	}
	e.JSON(http.StatusOK, api.Success(update))
	return nil
}
