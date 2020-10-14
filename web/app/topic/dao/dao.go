package dao

import (
	"github.com/Seven4X/link/web/app/topic/model"
	"github.com/Seven4X/link/web/library/store/db"
	"github.com/xormplus/xorm"
)

// 这里考虑要不要把 db暴露出去，简单的操作直接在service中做，这样反模式了，跨层调用了
type Dao struct {
	engine *xorm.Engine
}

func New() (dao *Dao) {
	dao = &Dao{
		engine: db.NewDb(),
	}
	return
}

func (dao *Dao) Save(topic *model.Topic) (i int64, err error) {
	i, err = dao.engine.InsertOne(topic)
	return i, err
}
