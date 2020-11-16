package dao

import (
	"github.com/Seven4X/link/web/app/link/model"
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

func (dao *Dao) Save(link *model.Link) (i int, err error) {
	_, err = dao.Insert(link)
	return link.Id, err
}
