package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUse(t *testing.T) {

	filter := GetCuckooFilter()
	assert.NotNil(t, filter)
	filter.Insert([]byte("https://www.yuque.com/dashboard"))
	filter.Insert([]byte("dashboard"))
	filter.Insert([]byte("dashboard3"))
	DumpCuckooFilter()
	filter = GetCuckooFilter()
	b := filter.Lookup([]byte("dashboard3"))
	assert.True(t, b)

}
