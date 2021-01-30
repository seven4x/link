package comment

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"xorm.io/builder"
)

func TestDao_ListHotCommentByLinkId(t *testing.T) {
	dao := NewDao()
	res, err := dao.ListHotCommentByLinkId([]interface{}{1, 2, 3, 4})
	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.NotEmpty(t, res[0].Creator.Avatar)
}

func TestBuilderUser(t *testing.T) {

	str, res, err := builder.ToSQL(builder.In("c2.link_id", 1, 2, 3, 4))
	println(str)
	println(res)
	println(err)

}
