package db

import (
	"encoding/json"
	"github.com/seven4x/link/app"
	"time"
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
	ListTopic = `select a.id,a.name,a.short_code
					from topic a
							 left join topic_rel b on a.id = b.bid
					where b.aid= ?
					  and b.position= ?
					  and a.id > ?`
	// 上级 左
	RevertListTopic = `select a.id,a.name,a.short_code
						from topic a
								 left join topic_rel b on a.id = b.aid
						where b.bid= ?
						  and b.position= ?
						  and a.id > ?`

	ListAllRelTopic = `select a.id,a.name,a.short_code
					from topic a
							 left join topic_rel b on a.id = b.bid or a.id = b.aid 
					where (b.aid= ? or b.bid = ?)
					  and a.id != ?`
)

func (dao *Dao) SaveTopic(topic *Topic, rel *TopicRel) (i int, err error) {
	//开启事物
	id, err := dao.Transaction(func(session *xorm.Session) (interface{}, error) {
		_, err := session.InsertOne(topic)
		if err != nil {
			session.Rollback()
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

		// 插入表单提交的关系
		rel.Position = convertPositionValue(rel.Position)
		// 如果添加相邻关系，同时需要添加一个上下关系
		if rel.Position == 2 {
			parent := new(TopicRel)
			_, err := session.SQL("select aid from topic_rel where bid=? and position=1", relId).Get(parent)
			notAllow := parent.Aid == 0 && topic.CreateBy != app.AdminId
			if !notAllow {
				relB := TopicRel{
					Aid:        parent.Aid,
					Bid:        id,
					Position:   1,
					CreateBy:   rel.CreateBy,
					Predicate:  "",
					CreateTime: rel.CreateTime,
				}
				if _, err = session.InsertOne(relB); err != nil {
					session.Rollback()
					return 0, err
				}
			}
		}

		if _, err = session.InsertOne(rel); err != nil {
			session.Rollback()
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

func (dao *Dao) GetByAlias(id string) (topic *Topic, err error) {
	alias := &TopicAlias{Alias: id}
	if _, err := dao.Get(alias); err != nil || alias.TopicId == 0 {
		return nil, err
	}
	topic = &Topic{Id: alias.TopicId}
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
	case 1, 2:
		return 1
	case 3, 4:
		return 2
	}
	return 1
}

/*ListRelativeTopic
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
		dbPosition = id
		sql = ListAllRelTopic //todo 优化性能
		prev = id
		break

	}
	topic = make([]Topic, 0)
	err = dao.SQL(sql, id, dbPosition, prev).Find(&topic)
	return
}

func (dao *Dao) SearchTopic(keyword string, prev int, size int) (res []Topic, hasMore bool, err error) {
	res = make([]Topic, 0)
	err = dao.Table("topic_shadow").Unscoped().
		Cols("name", "id", "short_code").
		Where("id>?", prev).
		And("name like ?", "%"+keyword+"%").Limit(size+1, 0).Find(&res)
	l := len(res)
	if l == 0 {
		return res, false, err
	}
	if l < size {
		return res, false, err
	}
	return res[:len(res)-1], len(res) > size, err

}

func (dao *Dao) ListHotTopic(limit int, start, end time.Time) (res []int, err error) {
	res = make([]int, 0)
	err = dao.SQL(`select topic_id
		from link
		where create_time between ? and ?
		group by topic_id limit ?`, start, end, limit).Find(&res)
	str, _ := json.Marshal(res)
	app.Infof("%s~%s,hot_topic:%s", start.String(), end.String(), str)
	return res, err
}
