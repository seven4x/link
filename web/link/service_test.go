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
	res, err := svr.ListLinkNoJoin(&req)
	assert.Nil(t, err)
	assert.NotNil(t, res)

}

//BenchmarkService_ListLink-8   	      42	  34524338 ns/op
func BenchmarkService_ListLink(b *testing.B) {
	svr := NewService()
	req := ListLinkRequest{
		Prev:   0,
		Size:   10,
		Tid:    1,
		UserId: 1,
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := svr.ListLinkNoJoin(&req)
		assert.Nil(b, err)
		assert.NotNil(b, res)
	}

}
