package job

import (
	"github.com/seven4x/link/db"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRefreshHotTopic(t *testing.T) {

	err := RefreshHotTopic()
	assert.Nil(t, err)
}

func TestInsertTime(t *testing.T) {

	dao := db.NewDao()

	now := time.Now()
	expireD, _ := time.ParseDuration("24h")
	expireTime := now.Add(expireD)
	ts := &db.HotTopic{
		Id:         99,
		Expire:     expireTime,
		CreateTime: now,
	}
	_, e := dao.Insert(ts)

	ts2 := &db.HotTopic{
		Id: 99,
	}
	dao.Find(ts2)
	assert.Nil(t, e)
}
