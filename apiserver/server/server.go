package server

import (
	"github.com/labstack/echo/v4"
	"github.com/seven4x/link/service"
)

const (
	DefaultSize = 10
)

type Server struct {
	svr  *service.Service
	echo *echo.Echo
}

func NewServer(e *echo.Echo) *Server {
	svr := service.NewService()
	return &Server{
		svr:  svr,
		echo: e,
	}
}

func (s *Server) InitRouter() {
	// 初始化模块
	s.RouterComment()
	s.RouterUser()
	s.RouterLink()
	s.RouterTopic()
	s.RouterVote()
}
