package comment

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDao_ListHotCommentByLinkId(t *testing.T) {
	dao := NewDao()
	res, err := dao.ListHotCommentByLinkId([]int{1, 2, 3})
	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.NotNil(t, res[0].CreatorAvatar)
}
