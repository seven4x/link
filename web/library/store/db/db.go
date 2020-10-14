package db

import (
	"github.com/Seven4X/link/web/library/log"
	_ "github.com/lib/pq"
	"github.com/xormplus/xorm"
)

var db *xorm.Engine

func init() {
	var err error
	//todo   viper
	db, err = xorm.NewPostgreSQL("postgres://postgres:linkhub2333@127.0.0.1:5432/link_hub?sslmode=disable")
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
	db.ShowSQL(true)
	err = db.Ping()
	if err != nil {
		log.Error("db-ping-failed:", err.Error())
	}
}

//调用时，p=bil
func NewDb() (p *xorm.Engine) {
	p = db
	return db
}
