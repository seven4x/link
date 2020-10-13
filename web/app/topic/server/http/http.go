/*
web层：route配置，参数解析，校验
*/
package http

import (
	"github.com/Seven4X/link/web/app/topic/api/request"
	"github.com/Seven4X/link/web/app/topic/model"
	"github.com/Seven4X/link/web/app/topic/service"
	"github.com/labstack/echo/v4"
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
func newTopic(c echo.Context) error {
	req := new(request.NewTopicReq)

	err := c.Bind(req)
	println(err.Error())
	//todo convert
	m := model.Topic{}
	svc.Save(&m)
	return nil
}
