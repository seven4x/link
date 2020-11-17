package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUse(t *testing.T) {

	filter := GetFilter()
	assert.NotNil(t, filter)

	//
	//filter.Insert([]byte("https://www.yuque.com/dashboard"))
	//filter.Insert([]byte("dashboard"))
	//filter.Insert([]byte("dashboard3"))

	b := filter.Lookup([]byte("dashboard3"))

	assert.True(t, b)

	DumpFilter()
}
