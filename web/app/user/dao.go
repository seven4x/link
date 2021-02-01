package user

import (
	"errors"
	"github.com/Seven4X/link/web/lib/store/db"
	"time"
	"xorm.io/xorm"
)

type Dao struct {
	*xorm.Engine
}

func (d *Dao) GetRegisterInfoByCode(code string) (*RegisterCode, error) {
	r := &RegisterCode{Code: code}
	d.Get(r)
	if r.Id == 0 {
		return nil, errors.New("error code")
	}
	return r, nil
}

/*
 1 新建用户
 2 创建注册记录
*/
func (d *Dao) Register(rc *RegisterCode, req *RegisterRequest) error {
	sess := d.NewSession()
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

func NewDao() (dao *Dao) {
	dao = &Dao{db.NewDb()}
	return
}
