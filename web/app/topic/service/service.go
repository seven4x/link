package service

import (
	"github.com/Seven4X/link/web/app/risk"
	"github.com/Seven4X/link/web/app/topic/dao"
	"github.com/Seven4X/link/web/app/topic/model"
	"github.com/Seven4X/link/web/library/log"
	"strconv"
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
2.检查关联topic是否存在 todo
3.检查是否重复
*/
func (service *Service) Save(topic *model.Topic) (rb bool, rs string) {
	var b, s = risk.IsAllowText(topic.Name)
	if !b {
		return false, s
	}

	i, err := service.dao.Save(topic)
	if err != nil {
		return false, err.Error()
	}

	log.Infow("save-new-topic", "id", i, "name", topic.Name)
	return true, strconv.FormatInt(i, 10)
}
