package service

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/seven4x/link/api"
	"github.com/seven4x/link/app"
	"github.com/seven4x/link/db"
	"math/rand"
	"strings"
)

// Login string 成功是jwt-token,失败是错误消息
func (s *Service) Login(l api.Login) (res *api.LoginResponse, err error) {
	if strings.ContainsAny(l.Username, "'<>@#&|") {
		return nil, errors.New("非法登陆")
	}
	u := db.Account{UserName: l.Username}
	if _, err := s.Dao.Get(&u); err != nil {
		return nil, err
	}
	if u.Id == 0 {
		return nil, errors.New("账户错误")
	}
	if u.Password != md5password(l.Password) {
		return nil, errors.New("密码错误")
	}
	token, claims, err := app.BuildToken(l.Username, u.Id)
	if err != nil {
		return nil, err
	}
	res = &api.LoginResponse{
		Token:    token,
		ExpireAt: claims.ExpiresAt,
	}
	return res, nil
}

func md5password(originPwd string) string {
	pwd := fmt.Sprintf("%x", md5.Sum([]byte(originPwd)))
	return pwd
}

func (s *Service) Register(req *api.RegisterRequest) (b bool, err error) {
	if ri, err := s.Dao.GetRegisterInfoByCode(req.Code); err != nil {
		return false, app.NewError(api.REGISTER_CODE_ERROR)
	} else {
		u := db.Account{UserName: req.LoginId}
		s.Dao.Get(&u)
		if u.Id != 0 {
			return false, app.NewError(api.REGISTER_NAME_REPEAT)
		}

		req.Password = md5password(req.Password)
		if err := s.Dao.Register(ri, req); err != nil {
			return false, app.NewError(err.Error())
		}

	}
	return true, nil
}

func (s *Service) GeneratorRegisterCode(id int) (string, error) {
	r := db.RegisterCode{UserId: id}
	s.Dao.Get(&r)
	if r.Code != "" {
		return r.Code, nil
	}
	rc := db.RegisterCode{
		UserId: id,
		Code:   RandStringBytes(8),
	}
	if _, err := s.Dao.InsertOne(&rc); err != nil {
		return "", err
	}
	return rc.Code, nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
