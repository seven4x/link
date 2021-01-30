package topic

import (
	"errors"
	"github.com/Seven4X/link/web/lib/store/db"
	"xorm.io/xorm"
)

const (
	FindByNameWithSameParent = ` select a.name
						from topic a
								 inner join topic_rel b
											on  (a.id = b.bid or a.id = b.aid)
						where b.position = ?
						  and (b.aid = ? or b.bid = ?)
						  and a.name = ?`
	//下级 右
	ListTopic = `select a.id,a.name
					from topic a
							 left join topic_rel b on a.id = b.bid
					where b.aid= ?
					  and b.position= ?
					  and a.id > ?`
	// 上级 左
	RevertListTopic = `select a.id,a.name
						from topic a
								 left join topic_rel b on a.id = b.aid
						where b.bid= ?
						  and b.position= ?
						  and a.id > ?`
)

// 这里考虑要不要把 db暴露出去，简单的操作直接在service中做，这样反模式了，跨层调用了
type Dao struct {
	*xorm.Engine
}

func NewDao() (dao *Dao) {
	dao = &Dao{
		Engine: db.NewDb(),
	}
	return
}

func (dao *Dao) Save(topic *Topic, rel *TopicRel) (i int, err error) {
	//开启事物
	id, err := dao.Transaction(func(session *xorm.Session) (interface{}, error) {
		_, err := session.InsertOne(topic)
		if err != nil {
			return 0, err
		}
		relId := rel.Aid
		id := topic.Id
		switch rel.Position {
		case 1:
			rel.Aid = id
			rel.Bid = relId
			break
		case 2:
			rel.Aid = relId
			rel.Bid = id
			break
		case 3:
			rel.Aid = id
			rel.Bid = relId
			break
		case 4:
			rel.Aid = relId
			rel.Bid = id
			break
		}
		rel.Position = convertPositionValue(rel.Position)

		if _, err = session.InsertOne(rel); err != nil {
			return 0, err
		}
		return id, nil
	})
	return id.(int), err
}

func (dao *Dao) ExistById(id int) (has bool, err error) {
	has, err = dao.Exist(&Topic{Id: id})
	return has, err
}

func (dao *Dao) GetById(id int) (topic *Topic, err error) {
	topic = &Topic{Id: id}
	b, err := dao.Get(topic)
	if b {
		return topic, nil
	} else {
		return nil, nil
	}
}

// 校验相同位置是否
func (dao *Dao) FindByNameWithSameParent(name string, position int, refId int) (b bool, err error) {
	topic := &Topic{}
	_, err = dao.SQL(FindByNameWithSameParent, convertPositionValue(position), refId, refId, name).Get(topic)
	return topic.Name != name, err
}

func convertPositionValue(position int) int {
	switch position {
	case 1:
	case 2:
		return 1
	case 3:
	case 4:
		return 2
	}
	return 1
}

/*
db 1 上下 2 左右
request 1,2,3,4 上下左右
*/
func (dao *Dao) ListRelativeTopic(id int, position string, prev int) (topic []Topic, err error) {
	var dbPosition int
	var sql string
	switch position {
	case "1":
		dbPosition = 1
		sql = RevertListTopic
		break
	case "2":
		dbPosition = 1
		sql = ListTopic
		break
	case "3":
		dbPosition = 2
		sql = RevertListTopic
		break
	case "4":
		dbPosition = 2
		sql = ListTopic
		break
	default:
		return nil, errors.New("position error，not in (1,2,3,4)")

	}
	topic = make([]Topic, 0)
	err = dao.SQL(sql, id, dbPosition, prev).Find(&topic)
	return
}

func (dao *Dao) SearchTopic(keyword string, prev int, size int) (res []Topic, hasMore bool, err error) {
	res = make([]Topic, 0)
	err = dao.Table("topic").
		Cols("name", "id").
		Where("id>?", prev).
		And("name like ?", "%"+keyword+"%").Limit(size+1, 0).Find(&res)
	if len(res) == 0 {
		return res, false, err
	}
	return res[:len(res)-1], len(res) > size, err

}
