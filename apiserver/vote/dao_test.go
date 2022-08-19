package vote

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDao_ListByBusinessId(t *testing.T) {
	dao := NewDao()
	res, err := dao.ListUserVoteByBusinessId([]interface{}{1, 2, 3, 4, 5}, 1, "l")
	assert.Nil(t, err)
	assert.NotNil(t, res)

}
