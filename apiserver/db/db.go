package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/seven4x/link/app"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var engine *xorm.Engine

func init() {
	var err error

	engine, err = xorm.NewEngine("sqlite3", buildSqlite3Dsn())
	if err != nil {
		app.Error(err.Error())
		panic(err)
	}
	engine.DatabaseTZ = time.Local
	engine.TZLocation = time.Local
	engine.ShowSQL(true)
	names.NewPrefixMapper(names.SnakeMapper{}, "t_")
	err = engine.Ping()
	if err != nil {
		app.Error("engine-ping-failed:", err.Error())
	}
}

func buildSqlite3Dsn() string {
	path := app.GetConfig("sqlite3.path")
	if path == "" {
		path = "~/.link"
	}
	return fmt.Sprintf("file:%s?cache=shared", path)

}

// NewDb 调用时，p=bil
func NewDb() (p *xorm.Engine) {
	p = engine
	return engine
}

// SetMockDb 参考：https://blog.marvel6.cn/2020/01/test-and-mock-db-by-xorm-with-the-help-of-convey-and-sqlmock/
func SetMockDb(mockDb *sql.DB) {
	engine.DB().DB = mockDb
}

func CloseEngine() {
	if engine != nil {
		_ = engine.Close()
	}
}
