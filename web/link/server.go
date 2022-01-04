package link

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/seven4x/link/web/messages"
	"strconv"

	"github.com/seven4x/link/web/middleware"
	"github.com/seven4x/link/web/util"
	"net/http"
)

var (
	svr = NewService()
)

/*Router
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
		//page next方式只能支持按照时间排序
		return listLink(e, func(request *ListLinkRequest) {
			request.OrderBy = "link.id desc"
		})
	}, middleware.Anonymous())
	g.GET("/marks/mine", func(e echo.Context) error {
		return listLink(e, func(request *ListLinkRequest) {
			request.FilterMy = true
		})
	}, middleware.JWT())
	g.GET("/marks/my", myAllPostLink, middleware.JWT())

	g.GET("/preview-token", getPreviewToken)

	g.GET("/recent", getRecentLinks)
}

func getRecentLinks(context echo.Context) error {
	prev := context.QueryParam("prev")
	prevI, _ := strconv.Atoi(prev)
	res, err := svr.GetRecentLinks(prevI)
	nextId := 0
	if len(res) > 0 {
		nextId = res[len(res)-1].Id
	}
	vos := make([]RecentVO, 0)
	for _, l := range res {
		v := RecentVO{l.Title, l.Id, l.Link, l.Tags, l.CreateAt}
		vos = append(vos, v)
	}
	context.JSON(200, struct {
		Data   []RecentVO
		NextId int
	}{vos, nextId})
	return err
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
		e.JSON(http.StatusBadRequest, util.Fail(err.Error()))
		return nil
	}
	if err := e.Validate(req); err != nil {
		e.JSON(http.StatusBadRequest, util.Fail(err.Error()))
		return nil
	}
	u := e.Get(util.User)
	if u == nil {
		e.JSON(http.StatusBadRequest, util.FailMsgId(messages.GlobalActionMustLogin))
		return nil
	}
	user := u.(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	link := req.ToLink()
	link.CreateBy = claims.Id
	id, err := svr.Save(link)
	if err != nil {
		_ = e.JSON(http.StatusInternalServerError, util.Response(id, err))
	} else {
		_ = e.JSON(http.StatusOK, util.Response(id, err))
	}

	return nil
}

func listLink(e echo.Context, setupRequest func(request *ListLinkRequest)) error {
	req := new(ListLinkRequest)
	err := e.Bind(req)
	if err != nil {
		return err
	}
	if err := e.Validate(req); err != nil {
		return err
	}
	uid := util.GetUserId(e)
	req.UserId = uid
	setupRequest(req)
	res, err := svr.ListLinkNoJoin(req)
	if err != nil {
		return err
	}
	l := len(res)
	hasMore := l == req.Size
	nextId := 0
	if l > 0 {
		nextId = res[len(res)-1].Id
	}
	e.Response().Header().Add("Cache-Control", "max-age=1800")
	return e.JSON(http.StatusOK, util.ResponsePage(res, hasMore, nextId))
}

func myAllPostLink(context echo.Context) error {
	return nil
}
