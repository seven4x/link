package service

import "github.com/seven4x/link/db"

type Service struct {
	Dao *db.Dao //todo fixme
}

func NewService(dao *db.Dao) *Service {
	return &Service{
		Dao: dao,
	}
}
