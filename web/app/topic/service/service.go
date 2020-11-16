package service

import (
	"github.com/Seven4X/link/web/app/risk"
	"github.com/Seven4X/link/web/app/topic/api/response"
	"github.com/Seven4X/link/web/app/topic/dao"
	"github.com/Seven4X/link/web/app/topic/model"
	"github.com/Seven4X/link/web/library/api"
	"github.com/Seven4X/link/web/library/api/messages"
	"github.com/Seven4X/link/web/library/log"
)

type Service struct {
	dao *dao.Dao
}

func NewService() (s *Service) {
	return &Service{
		dao: dao.New(),
	}
}

/*
1.敏感词过滤
2.检查关联topic是否存在
3.检查是否重复
*/
func (service *Service) Save(topic *model.Topic, rel *model.TopicRel) (id int, svrError *api.Err) {
	//todo 单用户创建频次限

	if topic.Lang == "zh" {
		var b, s = risk.IsAllowText(topic.Name)
		if !b {
			log.Infow("topic-save-not-allow", "keyword", s)
			return -1, api.NewError(messages.TopicContentNotAllowed)
		}
	}
	if rel.Aid == 0 {
		return -1, api.NewError(messages.TopicRootNotAllowed)
	}
	has, err := service.dao.ExistById(rel.Aid)
	if err != nil || !has {
		return -1, api.NewError(messages.TopicRefTopicNoExist)
	}
	has, err = service.dao.FindByNameWithSameParent(topic.Name, rel.Position, rel.Aid)
	if err != nil || !has {
		return -1, api.NewError(messages.TopicRepeatInSamePosition)
	}

	i, err := service.dao.Save(topic, rel)
	if err != nil {
		return -1, api.NewError(messages.TopicBackendDatabaseError)
	}
	log.Infow("save-new-topic", "uid", topic.CreateBy, "aid", rel.Aid, "name", topic.Name)
	return i, nil
}

func (service *Service) GetDetail(id int) (detail *response.TopicDetail, err error) {
	topic, err := service.dao.GetById(id)
	detail = response.TopicDetailOfModel(topic)
	//todo 其他关联查询信息

	return
}

func (service *Service) ListRelativeTopic(id int, position string, prev int) (topic []*response.TopicSimple, e error) {
	if topics, err := service.dao.ListRelativeTopic(id, position, prev); err == nil {
		return response.ModelToTopicSimple(topics), nil
	} else {
		return nil, err
	}

}
