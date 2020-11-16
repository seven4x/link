package server

import (
	"github.com/Seven4X/link/web/app/link/server/request"
	"github.com/Seven4X/link/web/app/link/service"
	"github.com/Seven4X/link/web/library/api"
	"github.com/Seven4X/link/web/library/api/messages"
	"github.com/Seven4X/link/web/library/config"
	"github.com/Seven4X/link/web/library/consts"
	"github.com/Seven4X/link/web/library/echo/mymw"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	svr = service.NewService()
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
	g.POST("", createLink, mymw.JWT())
	g.GET("", listLink)
	g.GET("/marks/hot", hotLink)
	g.GET("/marks/newest", newestLink)
	g.GET("/marks/mine", mineLink, mymw.JWT())
	g.GET("/marks/my", myAllPostLink, mymw.JWT())
	g.POST("/:lid/comment", postComment, mymw.JWT())
	g.GET("/:lid/comment", listComment)
	g.DELETE("/:lid/comment/:mid", deleteComment, mymw.JWT())
	g.GET("/preview-token", getPreviewToken)
}

func getPreviewToken(e echo.Context) error {
	str := config.GetString(config.LinkPreviewToken)

	_ = e.HTML(http.StatusOK, str)
	return nil
}

func createLink(e echo.Context) error {
	req := new(request.CreateLinkRequest)
	if err := e.Bind(req); err != nil {
		e.JSON(http.StatusOK, api.Fail(err.Error()))
		return nil
	}
	if err := e.Validate(req); err != nil {
		e.JSON(http.StatusOK, api.Fail(err.Error()))
		return nil
	}
	u := e.Get(consts.User)
	if u == nil {
		e.JSON(http.StatusOK, api.FailMsgId(messages.GlobalActionMustLogin))
		return nil
	}
	user := u.(*jwt.Token)
	claims := user.Claims.(*mymw.JwtCustomClaims)
	link := req.ToLink()
	link.CreateBy = claims.Id
	id, err := svr.Save(link)
	_ = e.JSON(http.StatusOK, api.Response(id, err))

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
