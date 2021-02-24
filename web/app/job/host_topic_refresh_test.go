package job

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRefreshHotTopic(t *testing.T) {

	err := RefreshHotTopic()
	assert.Nil(t, err)
}
