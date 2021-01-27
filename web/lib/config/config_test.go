package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	s := Get("gopath")
	println(s)
	assert.Equal(t, s, "/Users/seven/go")

	get := func(key string) {
		println(key)
		println(viper.Get(key))
		println(viper.GetString(key))
		println("-----------")
	}

	get("acm_ak")
	get("acm_sk")

}

func TestOsLookUp(t *testing.T) {

	show := func(key string) {
		val, ok := os.LookupEnv(key)
		if !ok {
			fmt.Printf("%s not set\n", key)
		} else {
			fmt.Printf("%s=%s\n", key, val)
		}
	}

	show("USER")
	show("GOPATH")
	show("LINK_DSN")
	show("acm_ak")
	show("acm_sk")

}

func TestGetFromAcm(t *testing.T) {
	res := getAcmConfig("token")
	assert.NotEmpty(t, res)
}

func TestGetAcmInit(t *testing.T) {
	res := GetString("link-preview-token")
	assert.NotEmpty(t, res)
}
