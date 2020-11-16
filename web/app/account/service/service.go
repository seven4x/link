package service

import (
	"errors"
	"github.com/Seven4X/link/web/app/account/api/request"
	"github.com/Seven4X/link/web/app/account/api/response"
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
func (svr *Service) Login(l request.Login) (res *response.LoginResponse, err error) {
	if l.Username == "test" {
		token, claims := mymw.BuildToken(l.Username, 12)
		res = &response.LoginResponse{
			Token:    token,
			ExpireAt: claims.ExpiresAt,
		}
		return res, nil
	}

	return nil, errors.New("账号错误")
}
