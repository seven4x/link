package db

import (
	"database/sql"
	"github.com/Seven4X/link/web/library/log"
	_ "github.com/lib/pq"
	"github.com/xormplus/xorm"
)

var engine *xorm.Engine

func init() {
	var err error
	//todo   viper
	engine, err = xorm.NewPostgreSQL("postgres://postgres:linkhub2333@127.0.0.1:5432/link_hub?sslmode=disable")
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
	engine.ShowSQL(true)
	err = engine.Ping()
	if err != nil {
		log.Error("engine-ping-failed:", err.Error())
	}
}

//调用时，p=bil
func NewDb() (p *xorm.Engine) {
	p = engine
	return engine
}

//参考：https://blog.marvel6.cn/2020/01/test-and-mock-db-by-xorm-with-the-help-of-convey-and-sqlmock/
func SetMockDb(mockDb *sql.DB) {
	engine.DB().DB = mockDb
}
