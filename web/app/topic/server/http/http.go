/*
web层：route配置，参数解析，校验
*/
package http

import (
	"github.com/Seven4X/link/web/app/topic/api/request"
	"github.com/Seven4X/link/web/app/topic/service"
	"github.com/Seven4X/link/web/library/api"
	"github.com/Seven4X/link/web/library/log"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	svc = service.NewService()
)

// call by cmd
func Router(e *echo.Echo) {

	route(e)
}

func route(echo *echo.Echo) {
	echo.POST("/topic", newTopic)

}

/*
1.敏感词校验
2.重复校验
3.
*/
func newTopic(e echo.Context) error {
	req := new(request.NewTopicReq)
	//1.解析
	if err := e.Bind(req); err != nil {
		_ = e.JSON(http.StatusOK, api.Fail(err.Error()))
		log.Info(err.Error())
		return nil
	}
	//2.校验
	if err := e.Validate(req); err != nil {
		//todo 失败消息的可读性
		_ = e.JSON(http.StatusOK, api.Fail(err.Error()))
		return nil
	}

	//3.转换
	//复杂对象使用gconv
	//topic := &model.Topic{}
	//_ = gconv.Struct(req, topic)
	//简单对象在Request对象中定义转化方法
	topic := req.ToTopic()
	log.Info(topic)
	//4.处理
	b, s := svc.Save(topic)
	_ = e.JSON(http.StatusOK, api.Response(b, s))
	return nil
}
