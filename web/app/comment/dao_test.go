package comment

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDao_ListHotCommentByLinkId(t *testing.T) {
	dao := NewDao()
	res, err := dao.ListHotCommentByLinkId([]interface{}{1, 2, 3})
	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.NotEmpty(t, res[0].Creator.Avatar)
}
