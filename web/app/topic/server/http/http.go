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
	err := e.Bind(req)
	if err != nil {
		_ = e.JSON(http.StatusOK, api.Fail(err.Error()))
		log.Info(err.Error())
		return nil
	}
	e.JSON(http.StatusOK, api.Succ(1))

	return nil

	////todo convert
	//m := model.Topic{}
	//svc.Save(&m)
	//return nil
}
