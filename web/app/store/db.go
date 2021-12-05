package store

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-txdb"
	"github.com/Seven4X/link/web/app/util"
	_ "github.com/mattn/go-sqlite3"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var engine *xorm.Engine

func init() {
	var err error

	engine, err = xorm.NewEngine("sqlite3", buildSqlite3Dsn())
	if err != nil {
		util.Error(err.Error())
		panic(err)
	}
	engine.DatabaseTZ = time.Local
	engine.TZLocation = time.Local
	engine.ShowSQL(true)
	names.NewPrefixMapper(names.SnakeMapper{}, "t_")
	err = engine.Ping()
	if err != nil {
		util.Error("engine-ping-failed:", err.Error())
	}
}

func buildSqlite3Dsn() string {
	path := util.Get("sqlite3.path")

	return fmt.Sprintf("file:%s?cache=shared", path)

}

func buildDsn() string {
	name := util.Get("db.user")
	password := util.Get("db.password")
	host := util.Get("db.host")
	db := util.Get("db.database")
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
		util.Error(err)
	}
	engine.DB().DB = db
}

func CloseEngine() {
	if engine != nil {
		_ = engine.Close()
	}
}
