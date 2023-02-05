package service

import "github.com/seven4x/link/db"

type Service struct {
	Dao *db.Dao //todo fixme
}

func NewService() *Service {
	dao := db.NewDao()
	return &Service{
		Dao: dao,
	}
}
