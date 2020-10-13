package service

import (
	"github.com/Seven4X/link/web/app/risk"
	"github.com/Seven4X/link/web/app/topic/dao"
	"github.com/Seven4X/link/web/app/topic/model"
)

type Service struct {
	dao *dao.Dao
}

func NewService() (s *Service) {
	return &Service{
		dao: dao.New(),
	}
}

func (service *Service) Save(topic *model.Topic) (rb bool, rs string) {
	var b, s = risk.IsAllowText(topic.Name)
	if b {
		_, _ = service.dao.Save(topic)

		return true, ""
	} else {
		return false, s
	}

}
