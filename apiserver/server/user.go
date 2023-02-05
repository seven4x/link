package server

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/seven4x/link/api"
	"github.com/seven4x/link/app"
	"github.com/seven4x/link/app/log"
	"github.com/seven4x/link/db"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	appid  = "appid123"
	secret = "secret"
)

func (s *Server) RouterUser() {
	e := s.echo
	e.GET("/api1/wx/cb", s.WechatCallback)
	g := e.Group("/api1/account")
	g.POST("/login", s.Login)
	g.GET("/logout", s.Logout)
	g.POST("/register", s.Register)
	g.GET("/get-my-code", s.GeneratorRegisterCode, app.JWT())
	g.GET("/info", s.Info, app.JWT())

}
func (s *Server) GeneratorRegisterCode(e echo.Context) error {
	user := e.Get("user").(*jwt.Token)
	if user == nil {
		e.JSON(http.StatusOK, api.Fail("need login"))
		return nil
	}
	claims := user.Claims.(*app.JwtCustomClaims)

	if code, err := s.svr.GeneratorRegisterCode(claims.Id); err != nil {
		return e.JSON(http.StatusInternalServerError, api.Fail(err.Error()))
	} else {
		return e.JSON(http.StatusOK, api.Success(code))
	}
	return nil
}

func (s *Server) Register(e echo.Context) error {
	req := new(api.RegisterRequest)
	e.Bind(req)
	if err := e.Validate(req); err != nil {
		return err
	}
	res, err := s.svr.Register(req)
	return e.JSON(http.StatusOK, api.Response(res, err))
}

func (s *Server) Login(e echo.Context) error {
	req := &api.Login{}
	if err := e.Bind(req); err != nil {
		e.JSON(http.StatusOK, api.Fail(err.Error()))
		return nil
	}

	if data, err := s.svr.Login(*req); err == nil {
		cookie := new(http.Cookie)
		cookie.Name = "token"
		cookie.Value = data.Token
		cookie.Expires = time.Unix(data.ExpireAt, 0)
		cookie.Path = "/"
		e.SetCookie(cookie)
		e.JSON(http.StatusOK, api.Success(data))
	} else {
		e.JSON(http.StatusOK, api.Fail(err.Error()))
	}

	return nil
}

func (s *Server) Logout(e echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Path = "/"
	cookie.MaxAge = -99

	e.SetCookie(cookie)
	return e.JSON(http.StatusOK, api.Success(true))
}

//https://developers.weixin.qq.com/doc/oplatform/Website_App/WeChat_Login/Wechat_Login.html
func (s *Server) WechatCallback(e echo.Context) error {
	code := e.QueryParam("code")
	if code == "" {
		log.Error("wechatCallback token is empty")
		return nil
	}
	url := "https://common.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
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
		url = "https://common.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s"
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
		var info api.WechatUserInfo
		json.Unmarshal(body, &info)
		//todo

	}()

	return nil
}

func (s *Server) Info(e echo.Context) error {
	u := app.GetUser(e)
	log.Info(u)
	acc := db.Account{Id: u.Id}
	s.svr.Dao.Get(&acc)
	info := api.AccountInfo{
		Id:       u.Id,
		Name:     u.Name,
		NickName: acc.NickName,
		Avatar:   acc.Avatar,
	}
	e.JSON(http.StatusOK, api.Success(info))
	return nil
}
