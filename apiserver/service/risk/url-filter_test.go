package risk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsAllowUrl(t *testing.T) {
	var b bool
	b = IsAllowUrl("angola.org")
	assert.False(t, b)
	b = IsAllowUrl("http://bilibili.com")
	assert.True(t, b)
	b = IsAllowUrl("http://viu.tv/ch/")
	assert.False(t, b)
}

func Test_IsUrl(t *testing.T) {
	var b bool
	b, _ = IsUrl("zkaip.com")
	assert.True(t, b)
	b, _ = IsUrl("http://google.com")
	assert.True(t, b)
	b, _ = IsUrl("http://w.com/cn")
	assert.True(t, b)
	b, _ = IsUrl("http://192.158.0.1:90")
	assert.True(t, b)
	b, _ = IsUrl("http://w")
	assert.False(t, b)
	b, _ = IsUrl("fsw")
	assert.False(t, b)
	b, _ = IsUrl("http://192.158.1/1")
	assert.False(t, b)

}
