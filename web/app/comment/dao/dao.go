package dao

import (
	"github.com/Seven4X/link/web/library/store/db"
	"github.com/xormplus/xorm"
)

type Dao struct {
	*xorm.Engine
}

func New() (dao *Dao) {
	dao = &Dao{db.NewDb()}
	return
}
