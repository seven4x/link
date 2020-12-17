package link

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_ListLink(t *testing.T) {
	svr := NewService()
	req := ListLinkRequest{
		Prev:   0,
		Size:   10,
		Tid:    1,
		UserId: 1,
	}
	res, err := svr.listLinkNoJoin(&req)
	assert.Nil(t, err)
	assert.NotNil(t, res)

}

func TestService_listLinkJoin(t *testing.T) {
	svr := NewService()
	req := ListLinkRequest{
		Prev:   0,
		Size:   10,
		Tid:    1,
		UserId: 1,
	}
	res, err := svr.listLinkJoin(&req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
