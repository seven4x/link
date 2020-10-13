package db

import (
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
