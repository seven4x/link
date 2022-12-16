package db

import (
	"errors"
	"github.com/seven4x/link/api"
	"time"
)

func (dao *Dao) GetRegisterInfoByCode(code string) (*RegisterCode, error) {
	r := &RegisterCode{Code: code}
	dao.Get(r)
	if r.Id == 0 {
		return nil, errors.New("error code")
	}
	return r, nil
}

/*
 1 新建用户
 2 创建注册记录
*/
func (dao *Dao) Register(rc *RegisterCode, req *api.RegisterRequest) error {
	sess := dao.NewSession()
	defer sess.Close()
	sess.Begin()
	u := Account{
		UserName: req.LoginId,
		NickName: req.NickName,
		Password: req.Password,
		Avatar:   "",
	}
	if _, err := sess.InsertOne(&u); err != nil {
		sess.Rollback()
		return err
	}

	ri := RegisterInfo{
		Code:         req.Code,
		CreateBy:     rc.UserId,
		UsedBy:       u.Id,
		UsedUserName: req.LoginId,
		UsedTime:     time.Now(),
	}
	if _, err := sess.InsertOne(&ri); err != nil {
		sess.Rollback()
		return err
	}
	sess.Commit()

	return nil

}
