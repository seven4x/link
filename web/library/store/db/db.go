package db

import (
	"database/sql"
	"github.com/DATA-DOG/go-txdb"
	"github.com/Seven4X/link/web/library/log"
	_ "github.com/lib/pq"
	"github.com/xormplus/xorm"
)

var engine *xorm.Engine

func init() {
	var err error
	//todo   viper
	engine, err = xorm.NewPostgreSQL("postgres://roach:Q7gc8rEdS@127.0.0.1:26257/link_hub")
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

//参考：https://github.com/DATA-DOG/Go-txdb
func RegisterMockDriver() {
	txdb.Register("txdb", "postgres", "postgres://roach:Q7gc8rEdS@127.0.0.1:26257/link_hub")
	db, err := sql.Open("txdb", "identifier")
	if err != nil {
		log.Error(err)
	}
	engine.DB().DB = db
}

func CloseEngine() {
	if engine != nil {
		_ = engine.Close()
	}
}
