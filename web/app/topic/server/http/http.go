/*
web层：route配置，参数解析，校验
*/
package http

import (
	"github.com/Seven4X/link/web/app/topic/api/request"
	"github.com/Seven4X/link/web/app/topic/service"
	"github.com/Seven4X/link/web/library/api"
	"github.com/Seven4X/link/web/library/consts"
	"github.com/Seven4X/link/web/library/echo/mymw"
	"github.com/Seven4X/link/web/library/log"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	svc = service.NewService()
)

// call by cmd
func Router(e *echo.Echo) {
	g := e.Group("/link")
	g.POST("/topic", createTopic)
	g.GET("/topic", searchTopic)
	g.GET("/topic/:id", topicDetail)
	g.GET("/topic/marks/hot", hotTopic)
	g.GET("/topic/marks/random", randomTopic)
	g.GET("/topic/marks/recent", recentTopic)
	g.GET("/topic/:tid/related/:position", relativeTopic)

}

/*
1.敏感词校验
2.重复校验
3.
*/
func createTopic(e echo.Context) error {
	req := new(request.NewTopicReq)
	//1.解析
	if err := e.Bind(req); err != nil {
		_ = e.JSON(http.StatusOK, api.Fail(err.Error()))
		log.Info(err.Error())
		return nil
	}
	//2.校验
	if err := e.Validate(req); err != nil {
		_ = e.JSON(http.StatusOK, api.Fail(err.Error()))
		return nil
	}

	//3.转换
	//复杂对象使用gconv
	//topic := &model.Topic{}
	//_ = gconv.Struct(req, topic)
	//简单对象在Request对象中定义转化方法
	topic, rel := req.ToTopic()

	user := e.Get(consts.User).(*jwt.Token)
	claims := user.Claims.(*mymw.JwtCustomClaims)
	topic.CreateBy = claims.Id
	//4.处理
	id, svrErr := svc.Save(topic, rel)
	_ = e.JSON(http.StatusOK, api.Response(id, svrErr))
	return nil
}

func searchTopic(e echo.Context) error {
	return nil
}
func topicDetail(e echo.Context) error {
	return nil
}
func hotTopic(e echo.Context) error {
	return nil
}
func randomTopic(e echo.Context) error {
	return nil
}
func recentTopic(e echo.Context) error {
	return nil
}
func relativeTopic(e echo.Context) error {
	return nil
}
