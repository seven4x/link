package db

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-txdb"
	"github.com/Seven4X/link/web/lib/config"
	"github.com/Seven4X/link/web/lib/log"
	_ "github.com/lib/pq"
	"github.com/xormplus/xorm"
	"github.com/xormplus/xorm/names"
)

var engine *xorm.Engine

func init() {
	var err error
	name := config.Get("APP_DB_NAME")
	password := config.Get("APP_DB_PASSWORD")
	host := config.Get("APP_DB_HOST")
	db := config.Get("APP_DB_DATABASE")
	sdn := fmt.Sprintf("postgres://%s:%s@%s/%s", name, password, host, db)
	engine, err = xorm.NewPostgreSQL(sdn)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
	engine.ShowSQL(true)
	names.NewPrefixMapper(names.SnakeMapper{}, "t_")
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
