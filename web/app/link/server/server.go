package server

import (
	"github.com/Seven4X/link/web/library/config"
	"github.com/labstack/echo/v4"
	"net/http"
)

/*
流控配置
[{
    "resource":"GET:/link/preview-token",
    "metricType":0,
    "tokenCalculateStrategy":0,
    "controlBehavior":0,
    "count": 3
}]

验证明令： ab -n1000 -c100 http://localhost:1323/link/preview-token
*/
func Router(e *echo.Echo) {
	g := e.Group("/link")
	g.GET("/preview-token", getPreviewToken)
}

func getPreviewToken(e echo.Context) error {
	str := config.GetString(config.LinkPreviewToken)

	_ = e.HTML(http.StatusOK, str)
	return nil
}
