/*
参考：
https://github.com/spf13/viper
https://help.aliyun.com/document_detail/130146.html?spm=5176.acm.0.dexternal.28824a9bRqym6y
*/
package config

import (
	"fmt"
	"github.com/Seven4X/link/web/library/log"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
)

const (
	endpoint         = "acm.aliyun.com"
	namespaceId      = "16c0d799-8364-4011-9bd9-94556fe35fdb"
	LinkPreviewToken = "link-preview-token"
)

func init() {
	viper.AutomaticEnv()
	InitAcm()
}

var client config_client.IConfigClient

/*
idea存在获取不到环境变量的问题，解决办法参考：
将idea启动命令加入path，通过控制台启动/Applications/IntelliJ IDEA.app/Contents/MacOS
viper会进行转换，环境变量必须是大写才能取到
*/
func InitAcm() {
	ak := GetString("acm_ak")
	sk := GetString("acm_sk")
	if ak == "" || sk == "" {
		log.Info("acm.ak acm.sk not config")
		return
	}
	clientConfig := constant.ClientConfig{
		//
		Endpoint:       endpoint + ":8080",
		NamespaceId:    namespaceId,
		AccessKey:      ak,
		SecretKey:      sk,
		TimeoutMs:      5 * 1000,
		ListenInterval: 30 * 1000,
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig": clientConfig,
	})
	if err != nil {
		log.Error("acm client create error", err.Error())
	}
	if configClient == nil {
		log.Error("acm client configClient is null")
		return
	}
	client = configClient
	initLinkPreviewToken()

}

func initLinkPreviewToken() {
	err := client.ListenConfig(vo.ConfigParam{
		DataId: "token",
		Group:  "DEFAULT_GROUP",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("ListenConfig group:" + group + ", dataId:" + dataId + ", data:" + data)
			viper.Set(LinkPreviewToken, data)
		},
	})
	if err != nil {
		log.Error("acm client Listen error", err.Error())
	}
	res := getAcm(LinkPreviewToken)
	viper.Set(LinkPreviewToken, res)
}

func Get(key string) interface{} {
	return viper.Get(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func getAcm(key string) string {
	if client == nil {
		return ""
	}
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: key,
		Group:  "DEFAULT_GROUP"})

	if err != nil {
		log.Error("getAcm Error", err.Error())
	}
	log.Info(content)
	return content
}
