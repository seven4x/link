package db

import (
	"github.com/Seven4X/link/web/app/topic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Sync(t *testing.T) {
	DB.Sync2(topic.Topic{})

	info, err := DB.TableInfo(topic.Topic{})
	if err != nil {
		assert.Fail(t, "failed")
	}
	assert.NotNil(t, info)
}
