package db

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-txdb"
	"github.com/Seven4X/link/web/lib/config"
	"github.com/Seven4X/link/web/lib/log"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var engine *xorm.Engine

func init() {
	var err error

	engine, err = xorm.NewEngine("postgres", buildDsn())
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

func buildDsn() string {
	name := config.Get("db.user")
	password := config.Get("db.password")
	host := config.Get("db.host")
	db := config.Get("db.database")
	sdn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", name, password, host, db)
	return sdn
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
	txdb.Register("txdb", "postgres", buildDsn())
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
