package service

import (
	"github.com/Seven4X/link/web/app/account/api/request"
	"github.com/Seven4X/link/web/library/echo/mymw"
)

type Service struct {
}

func New() (s *Service) {
	s = &Service{}
	return s
}

/*
string 成功是jwt-token,失败是错误消息
*/
func (svr *Service) Login(l request.Login) (bool, string) {
	if l.Username == "test" {
		//todo
		return true, mymw.BuildToken(l.Username, 12)
	}

	return false, "非法登陆"
}
