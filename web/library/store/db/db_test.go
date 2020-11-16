package db

import (
	model2 "github.com/Seven4X/link/web/app/link/model"
	"github.com/Seven4X/link/web/app/topic/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Sync(t *testing.T) {
	db := NewDb()
	db.Sync2(model.Topic{})
	info, err := db.TableInfo(model.Topic{})
	if err != nil {
		assert.Fail(t, "failed")
	}
	assert.NotNil(t, info)
}

func Test_SyncLink(t *testing.T) {
	db := NewDb()
	err := db.Sync2(model2.Link{})
	assert.Nil(t, err)
}
