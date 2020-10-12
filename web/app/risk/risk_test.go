package risk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_AllowText(t *testing.T)  {
	b, s := IsAllowText("干尼玛")
	println(s)
	assert.False(t,b,"期望非法")

}
