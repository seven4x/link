package db

import (
	"github.com/Seven4X/link/web/app/log"
	_ "github.com/lib/pq"
	"github.com/xormplus/xorm"
)

var DB *xorm.Engine

func init() {
	var err error
	//todo 有没有类似Spring的外部化配置
	DB, err = xorm.NewPostgreSQL("postgres://postgres:linkhub2333@127.0.0.1:5432/link_hub?sslmode=disable")
	if err != nil {
		log.Error(err.Error())
	}
	DB.ShowSQL(true)
	DB.Ping()
}
