package user

import (
	"encoding/json"
	"fmt"
	"github.com/Seven4X/link/web/library/api"
	"github.com/Seven4X/link/web/library/echo/mymw"
	"github.com/Seven4X/link/web/library/log"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

var (
	svr    = NewService()
	appid  = "appid123"
	secret = "secret"
)

func Router(e *echo.Echo) {
	e.GET("/wx/cb", wechatCallback)
	g := e.Group("/account")
	g.POST("/login", login)
	g.POST("/register", register)
	g.GET("/get-my-code", generatorRegisterCode, mymw.JWT())
	g.GET("/info", info, mymw.JWT())
}

func generatorRegisterCode(e echo.Context) error {
	user := e.Get("user").(*jwt.Token)
	if user == nil {
		e.JSON(http.StatusOK, api.Fail("need login"))
		return nil
	}
	claims := user.Claims.(*mymw.JwtCustomClaims)

	if code, err := svr.GeneratorRegisterCode(claims.Id); err != nil {
		return e.JSON(http.StatusInternalServerError, api.Fail(err.Error()))
	} else {
		return e.JSON(http.StatusOK, api.Success(code))
	}
	return nil
}

func register(e echo.Context) error {
	req := new(RegisterRequest)
	e.Bind(req)
	if err := e.Validate(req); err != nil {
		return err
	}
	res, err := svr.Register(req)
	return e.JSON(http.StatusOK, api.Response(res, err))
}

func login(e echo.Context) error {
	req := &Login{}
	if err := e.Bind(req); err != nil {
		e.JSON(http.StatusOK, api.Fail(err.Error()))
		return nil
	}

	if data, err := svr.Login(*req); err == nil {
		e.JSON(http.StatusOK, api.Success(data))
	} else {
		e.JSON(http.StatusOK, api.Fail(err.Error()))
	}

	return nil
}

//https://developers.weixin.qq.com/doc/oplatform/Website_App/WeChat_Login/Wechat_Login.html
func wechatCallback(e echo.Context) error {
	code := e.QueryParam("code")
	if code == "" {
		log.Error("wechatCallback token is empty")
		return nil
	}
	url := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	fmt.Printf(url, appid, secret, code)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Error(err.Error())
		log.Error("wechatCallback oauth2 error")
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
	}
	/*
		{
		"access_token":"ACCESS_TOKEN",
		"expires_in":7200,
		"refresh_token":"REFRESH_TOKEN",
		"openid":"OPENID",
		"scope":"SCOPE",
		"unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL"
		}
	*/
	m := map[string]interface{}{}
	json.Unmarshal(body, &m)
	token, b := m["access_token"]
	openId, _ := m["openid"]
	if !b || openId == "" {
		log.Error("GetAccessToken")
		log.Info(string(body))
		return nil
	}
	e.JSON(http.StatusOK, "ok")
	go func() {
		url = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s"
		fmt.Printf(url, token, openId)
		infoResp, err := http.Get(url)
		defer infoResp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error(err.Error())
		}
		/*
			{
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

			}
		*/
		var info WechatUserInfo
		json.Unmarshal(body, &info)
		//todo

	}()

	return nil
}

func info(e echo.Context) error {
	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(*mymw.JwtCustomClaims)
	log.Info(claims)
	e.JSON(http.StatusOK, api.Success([2]interface{}{"鸡要文件", claims}))
	return nil
}
