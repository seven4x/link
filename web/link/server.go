package link

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/seven4x/link/web/messages"
	"github.com/seven4x/link/web/middleware"
	"github.com/seven4x/link/web/util"
	"net/http"
)

var (
	svr = NewService()
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
	g := e.Group("/api1/link")
	g.POST("", createLink, middleware.JWT())
	g.GET("", func(e echo.Context) error {
		return listLink(e, func(request *ListLinkRequest) {})
	}, middleware.Anonymous())
	g.POST("/actions/batch", batchImport, middleware.JWT())
	g.GET("/marks/hot", func(e echo.Context) error {
		return listLink(e, func(request *ListLinkRequest) {
			request.OrderBy = "link.agree desc"
		})
	}, middleware.Anonymous())
	g.GET("/marks/newest", func(e echo.Context) error {
		return listLink(e, func(request *ListLinkRequest) {
			request.OrderBy = "link.create_time desc"
		})
	}, middleware.Anonymous())
	g.GET("/marks/mine", func(e echo.Context) error {
		return listLink(e, func(request *ListLinkRequest) {
			request.FilterMy = true
		})
	}, middleware.JWT())
	g.GET("/marks/my", myAllPostLink, middleware.JWT())

	g.GET("/preview-token", getPreviewToken)
}

// 解析URL中所有 a标签中的超链接保存
func batchImport(e echo.Context) error {
	url := e.Param("url")

	return e.JSON(http.StatusOK, map[string]interface{}{
		"url":  url,
		"size": 0,
	})

}

func getPreviewToken(e echo.Context) error {
	str := util.GetString("link-preview-token")

	_ = e.HTML(http.StatusOK, str)
	return nil
}

func createLink(e echo.Context) error {
	req := new(CreateLinkRequest)
	if err := e.Bind(req); err != nil {
		e.JSON(http.StatusOK, util.Fail(err.Error()))
		return nil
	}
	if err := e.Validate(req); err != nil {
		e.JSON(http.StatusOK, util.Fail(err.Error()))
		return nil
	}
	u := e.Get(util.User)
	if u == nil {
		e.JSON(http.StatusOK, util.FailMsgId(messages.GlobalActionMustLogin))
		return nil
	}
	user := u.(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	link := req.ToLink()
	link.CreateBy = claims.Id
	id, err := svr.Save(link)
	_ = e.JSON(http.StatusOK, util.Response(id, err))

	return nil
}

func listLink(e echo.Context, setupRequest func(request *ListLinkRequest)) error {
	req := new(ListLinkRequest)
	e.Bind(req)
	if err := e.Validate(req); err != nil {
		return err
	}
	uid := util.GetUserId(e)
	req.UserId = uid
	setupRequest(req)
	res, total, err := svr.ListLink(req)
	data := util.ResponsePage(res, err, total, len(res) > 0)
	e.Response().Header().Add("Cache-Control", "max-age=1800")
	return e.JSON(http.StatusOK, data)
}

func myAllPostLink(context echo.Context) error {
	return nil
}