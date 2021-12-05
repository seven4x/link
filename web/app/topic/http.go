/*
web层：route配置，参数解析，校验
*/
package topic

import (
	"github.com/Seven4X/link/web/app/messages"
	"github.com/Seven4X/link/web/app/middleware"
	"github.com/Seven4X/link/web/app/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var (
	svc = NewService()
)

const (
	DefaultSize = 10
)

// call by cmd
func Router(e *echo.Echo) {
	g := e.Group("/api1/topic")
	g.POST("", createTopic, middleware.JWT())
	g.GET("", searchTopic)
	g.GET("/:id", topicDetail)
	g.GET("/marks/hot", hotTopic)
	g.GET("/marks/random", randomTopic)
	g.GET("/marks/recent", recentTopic)
	g.GET("/:tid/related/:position", relativeTopic)
	g.GET("/actions/delete/:id", removeTopic, middleware.JWT())
}

func removeTopic(e echo.Context) error {
	uid := util.GetUserId(e)
	if uid == 0 {
		return e.String(http.StatusBadRequest, "not allow")
	}
	id := e.Param("id")
	topic := new(Topic)
	res, err := svc.dao.ID(id).Unscoped().Delete(topic)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.String(http.StatusOK, strconv.Itoa(int(res)))
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
	req := new(CreateTopicRequest)
	//1.解析
	if err := e.Bind(req); err != nil {
		_ = e.JSON(http.StatusOK, util.Fail(err.Error()))
		util.Info(err.Error())
		return nil
	}
	//2.校验
	if err := e.Validate(req); err != nil {
		_ = e.JSON(http.StatusOK, util.Fail(err.Error()))
		return nil
	}

	//3.转换
	//复杂对象使用gconv
	//topic := &model.Topic{}
	//_ = gconv.Struct(req, topic)
	//简单对象在Request对象中定义转化方法
	topic, rel := req.ConvertRequestToTopicModel()
	u := e.Get(util.User)
	if u == nil {
		e.JSON(http.StatusOK, util.FailMsgId(messages.GlobalActionMustLogin))
		return nil
	}
	user := u.(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	topic.CreateBy = claims.Id
	//4.处理
	id, svrErr := svc.Save(topic, rel)
	_ = e.JSON(http.StatusOK, util.Response(id, svrErr))
	return nil
}

func searchTopic(e echo.Context) error {
	keyword := e.QueryParam("q")
	if keyword == "" {
		e.JSON(http.StatusBadRequest, util.FailMsgId(messages.GlobalParamWrong))
		return nil
	}
	prev := e.QueryParam("prev")
	size, _ := strconv.Atoi(e.QueryParam("size"))
	if size > DefaultSize || size == 0 {
		size = DefaultSize
	}
	prevInt, _ := strconv.Atoi(prev)
	res, hasMore, err := svc.SearchTopic(keyword, prevInt, size)
	if err != nil {
		util.Error(err.Error())
		return nil
	}
	return e.JSON(http.StatusOK, util.ResponseHasMore(res, hasMore))
}
func topicDetail(e echo.Context) error {
	id := e.Param("id")
	var topic *Detail
	if i, err := strconv.Atoi(id); err != nil {
		topic, _ = svc.GetDetailByAlias(id)
	} else {
		topic, _ = svc.GetDetailById(i)
	}
	if topic == nil {
		return e.JSON(http.StatusOK, util.FailMsgId(messages.TopicNotFound))
	}
	return e.JSON(http.StatusOK, util.Success(topic))
}

func hotTopic(e echo.Context) error {
	res, err := svc.ListHotTopic()
	if err != nil {
		return err
	}
	e.JSON(http.StatusOK, util.Response(res, nil))
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
		return e.JSON(http.StatusOK, util.Fail("param wrong"))
	} else {
		position := e.Param("position")
		prev := e.QueryParam("prev")
		prevInt := 0
		if prev != "" {
			prevInt, _ = strconv.Atoi(prev)
		}
		topics, err := svc.ListRelativeTopic(id, position, prevInt)
		if err == nil {
			return e.JSON(http.StatusOK, util.ResponseHasMore(topics, len(topics) > 0))
		} else {
			return err
		}
	}
}
