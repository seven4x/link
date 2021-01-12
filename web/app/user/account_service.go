package user

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/Seven4X/link/web/library/api"
	"github.com/Seven4X/link/web/library/api/messages"
	"github.com/Seven4X/link/web/library/echo/mymw"
	"math/rand"
	"strings"
)

type Service struct {
	dao *Dao
}

func NewService() (s *Service) {
	s = &Service{dao: NewDao()}
	return s
}

/*
string 成功是jwt-token,失败是错误消息
*/
func (svr *Service) Login(l Login) (res *LoginResponse, err error) {
	if strings.ContainsAny(l.Username, "'<>@#&|") {
		return nil, errors.New("非法登陆")
	}
	u := Account{UserName: l.Username}
	if _, err := svr.dao.Table("user").Get(&u); err != nil {
		return nil, err
	}
	if u.Id == 0 {
		return nil, errors.New("账户错误")
	}
	if u.Password != md5password(l.Password) {
		return nil, errors.New("密码错误")
	}
	token, claims, err := mymw.BuildToken(l.Username, u.Id)
	if err != nil {
		return nil, err
	}
	res = &LoginResponse{
		Token:    token,
		ExpireAt: claims.ExpiresAt,
	}
	return res, nil
}

func md5password(originPwd string) string {
	pwd := fmt.Sprintf("%x", md5.Sum([]byte(originPwd)))
	return pwd
}

func (svr *Service) Register(req *RegisterRequest) (b bool, err *api.Err) {
	if ri, err := svr.dao.GetRegisterInfoByCode(req.Code); err != nil {
		return false, api.NewError(messages.REGISTER_CODE_ERROR)
	} else {
		u := Account{UserName: req.LoginId}
		svr.dao.Get(&u)
		if u.Id != 0 {
			return false, api.NewError(messages.REGISTER_NAME_REPEAT)
		}

		req.Password = md5password(req.Password)
		if err := svr.dao.Register(ri, req); err != nil {
			return false, api.NewError(err.Error())
		}

	}
	return true, nil
}

func (svr *Service) GeneratorRegisterCode(id int) (string, error) {
	r := RegisterCode{UserId: id}
	svr.dao.Get(&r)
	if r.Code != "" {
		return r.Code, nil
	}
	rc := RegisterCode{
		UserId: id,
		Code:   RandStringBytes(8),
	}
	if _, err := svr.dao.InsertOne(&rc); err != nil {
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
