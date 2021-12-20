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
	res, total, err := dao.ListLink(req)
	assert.Nil(t, err)
	assert.True(t, total > 0)
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
