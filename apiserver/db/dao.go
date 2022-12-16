package db

import (
	"xorm.io/xorm"
)

type Dao struct {
	*xorm.Engine
}

func NewDao() (dao *Dao) {
	dao = &Dao{NewDb()}
	dao.NewSession()
	return
}
