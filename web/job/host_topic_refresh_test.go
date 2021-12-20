package job

import (
	"github.com/seven4x/link/web/topic"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRefreshHotTopic(t *testing.T) {

	err := RefreshHotTopic()
	assert.Nil(t, err)
}

func TestInsertTime(t *testing.T) {

	dao := topic.NewDao()

	now := time.Now()
	expireD, _ := time.ParseDuration("24h")
	expireTime := now.Add(expireD)
	ts := &topic.HotTopic{
		Id:         99,
		Expire:     expireTime,
		CreateTime: now,
	}
	_, e := dao.Insert(ts)

	ts2 := &topic.HotTopic{
		Id: 99,
	}
	dao.Find(ts2)
	assert.Nil(t, e)
}
