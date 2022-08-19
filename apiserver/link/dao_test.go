package link

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDao_ListLink(t *testing.T) {
	dao := NewDao()
	req := &ListLinkRequest{
		Prev:   0,
		Size:   10,
		Tid:    1,
		UserId: 1,
	}
	res, err := dao.ListLink(req)
	assert.Nil(t, err)
	assert.NotNil(t, res)

}

func TestCountLink(t *testing.T) {
	dao := NewDao()
	req := &ListLinkRequest{
		Prev:   0,
		Size:   10,
		Tid:    1,
		UserId: 1,
	}

	total, err := dao.countLink(req)

	assert.Nil(t, err)

	assert.True(t, total > 0)
}

func TestDao_GetRecentLinks(t *testing.T) {
	dao := NewDao()

	res, err := dao.GetRecentLinks(12)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, 10, len(res))
}
