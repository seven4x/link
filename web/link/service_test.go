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
	res, _, err := svr.listLinkNoJoin(&req)
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
	res, _, err := svr.listLinkJoin(&req)
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
		res, _, err := svr.listLinkNoJoin(&req)
		assert.Nil(b, err)
		assert.NotNil(b, res)
	}

}

//BenchmarkService_ListLinkJoin-8   	      42	  34399993 ns/op
// 更优，通过join减少了和服务器的网络交互，。数据量多了之后效果如何还不可知 todo
func BenchmarkService_ListLinkJoin(b *testing.B) {
	svr := NewService()
	req := ListLinkRequest{
		Prev:   0,
		Size:   10,
		Tid:    1,
		UserId: 1,
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, _, err := svr.listLinkJoin(&req)
		assert.Nil(b, err)
		assert.NotNil(b, res)
	}

}
