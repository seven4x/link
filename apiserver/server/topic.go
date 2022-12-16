package server

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/seven4x/link/api"
	"github.com/seven4x/link/app"
	"github.com/seven4x/link/db"
	"net/http"
	"strconv"
)

func RouterTopic(e *echo.Echo) {
	g := e.Group("/api1/topic")
	g.POST("", CreateTopic, app.JWT())
	g.GET("", SearchTopic)
	g.GET("/:id", TopicDetail)
	g.GET("/marks/hot", HotTopic)
	g.GET("/marks/random", RandomTopic)
	g.GET("/marks/recent", RecentTopic)
	g.GET("/:tid/related/:position", RelativeTopic)
	g.GET("/actions/delete/:id", RemoveTopic, app.JWT())
}
func RemoveTopic(e echo.Context) error {
	uid := app.GetUserId(e)
	if uid == 0 {
		return e.String(http.StatusBadRequest, "not allow")
	}
	id := e.Param("id")
	topic := new(db.Topic)
	res, err := svr.Dao.ID(id).Unscoped().Delete(topic)
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
func CreateTopic(e echo.Context) error {
	req := new(api.CreateTopicRequest)
	//1.解析
	if err := e.Bind(req); err != nil {
		_ = e.JSON(http.StatusOK, api.Fail(err.Error()))
		app.Info(err.Error())
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
	topic, rel := ConvertRequestToTopicModel(req)
	u := e.Get(app.User)
	if u == nil {
		e.JSON(http.StatusOK, api.FailMsgId(api.GlobalActionMustLogin))
		return nil
	}
	user := u.(*jwt.Token)
	claims := user.Claims.(*app.JwtCustomClaims)
	topic.CreateBy = claims.Id
	//4.处理
	id, svrErr := svr.SaveTopic(topic, rel)
	_ = e.JSON(http.StatusOK, api.Response(id, svrErr))
	return nil
}

func SearchTopic(e echo.Context) error {
	keyword := e.QueryParam("q")
	if keyword == "" {
		e.JSON(http.StatusBadRequest, api.FailMsgId(api.GlobalParamWrong))
		return nil
	}
	prev := e.QueryParam("prev")
	size, _ := strconv.Atoi(e.QueryParam("size"))
	if size > DefaultSize || size == 0 {
		size = DefaultSize
	}
	prevInt, _ := strconv.Atoi(prev)
	res, hasMore, err := svr.SearchTopic(keyword, prevInt, size)
	if err != nil {
		app.Error(err.Error())
		return nil
	}
	return e.JSON(http.StatusOK, api.ResponseHasMore(res, hasMore))
}
func TopicDetail(e echo.Context) error {
	id := e.Param("id")
	var topic *api.TopicDetail
	if i, err := strconv.Atoi(id); err != nil {
		topic, _ = svr.GetDetailByAlias(id)
	} else {
		topic, _ = svr.GetDetailById(i)
	}
	if topic == nil {
		return e.JSON(http.StatusOK, api.FailMsgId(api.TopicNotFound))
	}
	return e.JSON(http.StatusOK, api.Success(topic))
}

func HotTopic(e echo.Context) error {
	res, err := svr.ListHotTopic()
	if err != nil {
		return err
	}
	e.JSON(http.StatusOK, api.Response(res, nil))
	return nil
}
func RandomTopic(e echo.Context) error {
	return nil
}
func RecentTopic(e echo.Context) error {
	return nil
}
func RelativeTopic(e echo.Context) error {
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
		topics, err := svr.ListRelativeTopic(id, position, prevInt)
		if err == nil {
			return e.JSON(http.StatusOK, api.ResponseHasMore(topics, len(topics) > 0))
		} else {
			return err
		}
	}
}
