package user

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringFormat(t *testing.T) {
	url := "https://common.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	fmt.Printf(url, appid, secret, "abc")
	println(url)
	assert.NotNil(t, url)
}

func TestUnmarshal(t *testing.T) {
	/*

	 */
	m := map[string]interface{}{}
	json.Unmarshal([]byte(`{
		"access_token":"ACCESS_TOKEN",
		"expires_in":7200,
		"refresh_token":"REFRESH_TOKEN",
		"openid":"OPENID",
		"scope":"SCOPE",
		"unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL"
		}`), &m)

	assert.NotNil(t, m)
	var res WechatUserInfo
	json.Unmarshal([]byte(`{
"openid":"OPENID",
"nickname":"NICKNAME",
"sex":1,
"province":"PROVINCE",
"city":"CITY",
"country":"COUNTRY",
"headimgurl": "https://thirdwx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/0",
"privilege":[
"PRIVILEGE1",
"PRIVILEGE2"
],
"unionid": " o6_bmasdasdsad6_2sgVt7hMZOPfL"

}`), &res)

	assert.NotNil(t, res)
	assert.Equal(t, res.Nickname, "NICKNAME")
}
