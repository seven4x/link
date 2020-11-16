package service

import (
	"github.com/Seven4X/link/web/app/comment/dao"
	"github.com/Seven4X/link/web/app/comment/model"
)

type Service struct {
	dao *dao.Dao
}

func NewService() (service *Service) {
	service = &Service{
		dao: dao.New(),
	}
	return
}

func (service *Service) Save(comment *model.Comment) (id int, err error) {

	return -1, nil
}
