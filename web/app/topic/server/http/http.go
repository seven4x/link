/*
web层：route配置，参数解析，校验
*/
package http

import (
	"github.com/Seven4X/link/web/app/topic/api/request"
	"github.com/Seven4X/link/web/app/topic/service"
	"github.com/Seven4X/link/web/library/api"
	"github.com/Seven4X/link/web/library/api/messages"
	"github.com/Seven4X/link/web/library/consts"
	"github.com/Seven4X/link/web/library/echo/mymw"
	"github.com/Seven4X/link/web/library/log"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var (
	svc = service.NewService()
)

// call by cmd
func Router(e *echo.Echo) {
	g := e.Group("/topic")
	g.POST("", createTopic, mymw.JWT())
	g.GET("", searchTopic)
	g.GET("/:id", topicDetail)
	g.GET("/marks/hot", hotTopic)
	g.GET("/marks/random", randomTopic)
	g.GET("/marks/recent", recentTopic)
	g.GET("/:tid/related/:position", relativeTopic)

}

/*
1. 参数解析
2. 参数校验
3. 转换参数
4. 调用service
5.0 转换结果model到vo
5. JSON 响应

*/
func createTopic(e echo.Context) error {
	req := new(request.CreateTopicRequest)
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
	u := e.Get(consts.User)
	if u == nil {
		e.JSON(http.StatusOK, api.FailMsgId(messages.GlobalActionMustLogin))
		return nil
	}
	user := u.(*jwt.Token)
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
	id := e.Param("id")
	if i, err := strconv.Atoi(id); err != nil {
		_ = e.JSON(http.StatusOK, api.Fail("id must integer"))
	} else {
		topic, _ := svc.GetDetail(i)
		if topic == nil {
			return e.JSON(http.StatusOK, api.FailMsgId(messages.TopicNotFound))
		}
		return e.JSON(http.StatusOK, api.Success(topic))
	}
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
	tid := e.Param("tid")
	if id, err := strconv.Atoi(tid); err != nil {
		return e.JSON(http.StatusOK, api.Fail("param wrong"))
	} else {
		position := e.Param("position")
		prev := e.QueryParam("prev")
		prevInt := 0
		if prev != "" {
			prevInt, _ = strconv.Atoi(prev)
		}
		topics, err := svc.ListRelativeTopic(id, position, prevInt)
		if err == nil {
			return e.JSON(http.StatusOK, api.ResponseHasMore(topics, len(topics) > 0))
		} else {
			return err
		}
	}
}
