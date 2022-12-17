package db

import (
	"xorm.io/xorm"
)

type Dao struct {
	*xorm.Engine
}

func NewDao() (dao *Dao) {
	eng, err := NewDb()
	if err != nil {
		panic("couldn't create db ")
	}
	dao = &Dao{eng}
	return
}
