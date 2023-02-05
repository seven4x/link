package service

import (
	"github.com/seven4x/link/api"
	"github.com/seven4x/link/app"
	"github.com/seven4x/link/app/log"
	"github.com/seven4x/link/db"
	"github.com/seven4x/link/service/risk"
	"strconv"
	"time"
)

/*
1.敏感词过滤
2.检查关联topic是否存在
3.检查是否重复
*/
func (s *Service) SaveTopic(topic *db.Topic, rel *db.TopicRel) (id int, errs error) {

	if topic.Lang == "zh" {
		var b, s = risk.IsAllowText(topic.Name)
		if !b {
			log.Infow("topic-save-not-allow", "keyword", s)
			return -1, app.NewError(api.TopicContentNotAllowed)
		}
	}
	if rel.Aid == 0 {
		return -1, app.NewError(api.TopicRootNotAllowed)
	}
	has, err := s.Dao.ExistById(rel.Aid)
	if err != nil || !has {
		return -1, app.NewError(api.TopicRefTopicNoExist)
	}
	has, err = s.Dao.FindByNameWithSameParent(topic.Name, rel.Position, rel.Aid)
	if err != nil || !has {
		return -1, app.NewError(api.TopicRepeatInSamePosition)
	}

	i, err := s.Dao.SaveTopic(topic, rel)
	if err != nil {
		return -1, app.NewError(api.TopicBackendDatabaseError)
	}
	log.Infow("save-new-topic", "uid", topic.CreateBy, "aid", rel.Aid, "name", topic.Name)
	return i, nil
}

func (s *Service) GetDetailById(id int) (detail *api.TopicDetail, err error) {
	topic, err := s.Dao.GetById(id)
	detail = BuildDetailFromModel(topic)
	//todo 其他关联查询信息

	return
}
func (s *Service) GetDetailByAlias(id string) (detail *api.TopicDetail, err error) {
	topic, err := s.Dao.GetByAlias(id)
	detail = BuildDetailFromModel(topic)
	//todo 其他关联查询信息

	return
}
func (s *Service) ListRelativeTopic(id int, position string, prev int) (topic []*api.SnapShot, e error) {
	if topics, err := s.Dao.ListRelativeTopic(id, position, prev); err == nil {
		return ConvertModelToTopicSimple(topics), nil
	} else {
		return nil, err
	}

}

func (s *Service) SearchTopic(keyword string, prev int, size int) (res []*api.SnapShot, hasMore bool, err error) {
	topics, hasMore, err := s.Dao.SearchTopic(keyword, prev, size)
	res = make([]*api.SnapShot, 0)
	for _, t := range topics {
		res = append(res, &api.SnapShot{
			Name:      t.Name,
			Id:        strconv.Itoa(t.Id),
			ShortCode: t.ShortCode,
		})
	}
	return res, hasMore, err

}

// 每日统计写hot_topic表
func (s *Service) ListHotTopic() (res []api.SnapShot, err error) {
	topics := make([]api.SnapShot, 0)
	err = s.Dao.Table("hot_topic").
		Distinct("topic.name", "topic.id", "topic.short_code").
		Join("inner", "topic", "topic.id=hot_topic.id").
		Where("expire>?", time.Now()).
		Limit(10, 0).Find(&topics)
	return topics, err
}
