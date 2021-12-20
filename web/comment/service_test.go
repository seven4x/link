package comment

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_ListComment(t *testing.T) {
	svr := NewService()
	req := &ListCommentRequest{
		Prev:   0,
		SortBy: "newest",
		Size:   10,
		UserId: 1,
		LinkId: 1,
	}
	res, hasMore, err := svr.ListComment(req)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.True(t, hasMore)
}
