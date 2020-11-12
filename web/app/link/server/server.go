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
	g.POST("", createLink)
	g.GET("", listLink)
	g.GET("/marks/hot", hotLink)
	g.GET("/marks/newest", newestLink)
	g.GET("/marks/mine", mineLink)
	g.GET("/marks/my", myAllPostLink)
	g.POST("/:lid/comment", postComment)
	g.GET("/:lid/comment", listComment)
	g.DELETE("/:lid/comment/:mid", deleteComment)
	g.GET("/preview-token", getPreviewToken)
}

func getPreviewToken(e echo.Context) error {
	str := config.GetString(config.LinkPreviewToken)

	_ = e.HTML(http.StatusOK, str)
	return nil
}

func createLink(e echo.Context) error {

	return nil
}
func listLink(context echo.Context) error {
	return nil
}
func hotLink(context echo.Context) error {
	return nil
}
func newestLink(context echo.Context) error {
	return nil
}
func mineLink(context echo.Context) error {
	return nil
}
func myAllPostLink(context echo.Context) error {
	return nil
}
func postComment(context echo.Context) error {
	return nil
}
func listComment(context echo.Context) error {
	return nil
}
func deleteComment(context echo.Context) error {
	return nil
}
