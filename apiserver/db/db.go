package db

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/seven4x/link/app"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)


func initDB() (*xorm.Engine, error) {

	engine, err := xorm.NewEngine("sqlite3", buildSqlite3Dsn())
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
	return engine, err
}

func buildSqlite3Dsn() string {
	path := app.GetConfig("sqlite3.path")
	if path == "" {
		path = "~/.link"
	}
	return fmt.Sprintf("file:%s?cache=shared", path)

}

func NewDb() (*xorm.Engine, error) {
	return initDB()
}
