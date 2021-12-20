package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRegisterInfoByCode(t *testing.T) {
	dao := NewDao()
	coe, err := dao.GetRegisterInfoByCode("seven4x")
	assert.NotNil(t, coe)
	assert.Nil(t, err)
}
func TestGetNoFoundRegisterInfoByCode(t *testing.T) {
	dao := NewDao()
	coe, err := dao.GetRegisterInfoByCode("no-sax")
	assert.Nil(t, coe)
	assert.NotNil(t, err)
}
