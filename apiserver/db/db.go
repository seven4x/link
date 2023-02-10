package db

import (
	"fmt"
	"time"

	"github.com/seven4x/link/app"
	"github.com/seven4x/link/app/log"

	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

func NewDb() (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("sqlite3", fmt.Sprintf("file:%s?cache=shared", app.DbPath))
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
	engine.DatabaseTZ = time.Local
	engine.TZLocation = time.Local
	engine.ShowSQL(true)
	names.NewPrefixMapper(names.SnakeMapper{}, "t_")
	err = engine.Ping()
	if err != nil {
		log.Error("engine-ping-failed:", err.Error())
	}
	return engine, err
}
