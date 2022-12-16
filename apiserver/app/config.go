package app

/*
参考：
https://github.com/spf13/viper
https://help.aliyun.com/document_detail/130146.html?spm=5176.acm.0.dexternal.28824a9bRqym6y
*/

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.link")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		println(fmt.Errorf("Fatal error config file: %s \n", err))
	}

}

func GetConfig(key string) interface{} {
	return viper.Get(key)
}

func GetConfigString(key string) string {
	return viper.GetString(key)
}
