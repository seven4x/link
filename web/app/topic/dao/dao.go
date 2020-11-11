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

func (dao *Dao) Save(topic *model.Topic, rel *model.TopicRel) (i int64, err error) {
	//开启事物
	id, err := dao.engine.Transaction(func(session *xorm.Session) (interface{}, error) {
		if i, err = session.InsertOne(topic); err != nil {
			return nil, err
		}
		rel.Bid = i
		if _, err = session.InsertOne(rel); err != nil {
			return nil, err
		}
		return i, nil
	})
	return id.(int64), err
}

func (dao *Dao) ExistById(id int64) (has bool, err error) {
	has, err = dao.engine.Exist(&model.Topic{Id: id})
	return has, err
}

func (dao *Dao) GetById(id int64) (topic *model.Topic, err error) {
	topic = &model.Topic{Id: id}
	_, err = dao.engine.Get(topic)
	return
}

const (
	FindByNameWithSameParent = `select a.name from topic a inner join topic_rel b 
											on a.id = b.bid
											where b.position=? and b.aid=? and a.name=? `
)

// 校验相同位置是否
func (dao *Dao) FindByNameWithSameParent(name string, position int, refId int64) (b bool, err error) {
	topic := &model.Topic{}
	err = dao.engine.SQL(FindByNameWithSameParent, position, refId, name).Find(topic)
	return topic.Name == name, err
}
