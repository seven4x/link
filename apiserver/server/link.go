package server

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/seven4x/link/api"
	"github.com/seven4x/link/app"
	"net/http"
	"strconv"
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
func (s *Server) RouterLink() {
	e := s.echo
	g := e.Group("/api1/link")
	g.POST("", s.CreateLink, app.JWT())
	g.GET("", func(e echo.Context) error {
		return s.ListLink(e, func(request *api.ListLinkRequest) {})
	}, app.Anonymous())
	g.POST("/actions/batch", s.BatchImport, app.JWT())
	g.GET("/marks/hot", func(e echo.Context) error {
		return s.ListLink(e, func(request *api.ListLinkRequest) {
			request.OrderBy = "link.agree desc"
		})
	}, app.Anonymous())
	g.GET("/marks/newest", func(e echo.Context) error {
		//page next方式只能支持按照时间排序
		return s.ListLink(e, func(request *api.ListLinkRequest) {
			request.OrderBy = "link.id desc"
		})
	}, app.Anonymous())
	g.GET("/marks/mine", func(e echo.Context) error {
		return s.ListLink(e, func(request *api.ListLinkRequest) {
			request.FilterMy = true
		})
	}, app.JWT())
	g.GET("/marks/my", s.MyAllPostLink, app.JWT())

	g.GET("/preview-token", s.GetPreviewToken)

	g.GET("/recent", s.GetRecentLinks)
}

func (s *Server) GetRecentLinks(context echo.Context) error {
	prev := context.QueryParam("prev")
	prevI, _ := strconv.Atoi(prev)
	res, err := s.svr.GetRecentLinks(prevI)
	nextId := 0
	if len(res) > 0 {
		nextId = res[len(res)-1].Id
	}
	vos := make([]api.RecentVO, 0)
	for _, l := range res {
		v := api.RecentVO{l.Title, l.Id, l.Link, l.Tags, l.CreateAt}
		vos = append(vos, v)
	}
	context.JSON(200, struct {
		Data   []api.RecentVO
		NextId int
	}{vos, nextId})
	return err
}

// 解析URL中所有 a标签中的超链接保存
func (s *Server) BatchImport(e echo.Context) error {
	url := e.Param("url")

	return e.JSON(http.StatusOK, map[string]interface{}{
		"url":  url,
		"size": 0,
	})

}

func (s *Server) GetPreviewToken(e echo.Context) error {
	str := app.GetConfigString("link-preview-token")

	_ = e.HTML(http.StatusOK, str)
	return nil
}

func (s *Server) CreateLink(e echo.Context) error {
	req := new(api.CreateLinkRequest)
	if err := e.Bind(req); err != nil {
		e.JSON(http.StatusBadRequest, api.Fail(err.Error()))
		return nil
	}
	if err := e.Validate(req); err != nil {
		e.JSON(http.StatusBadRequest, api.Fail(err.Error()))
		return nil
	}
	u := e.Get(app.User)
	if u == nil {
		e.JSON(http.StatusBadRequest, api.FailMsgId(api.GlobalActionMustLogin))
		return nil
	}
	user := u.(*jwt.Token)
	claims := user.Claims.(*app.JwtCustomClaims)
	link := ToLink(req)
	link.CreateBy = claims.Id
	id, err := s.svr.SaveLink(link)
	if err != nil {
		_ = e.JSON(http.StatusInternalServerError, api.Response(id, err))
	} else {
		_ = e.JSON(http.StatusOK, api.Response(id, err))
	}

	return nil
}

func (s *Server) ListLink(e echo.Context, setupRequest func(request *api.ListLinkRequest)) error {
	req := new(api.ListLinkRequest)
	err := e.Bind(req)
	if err != nil {
		return err
	}
	if err := e.Validate(req); err != nil {
		return err
	}
	uid := app.GetUserId(e)
	req.UserId = uid
	setupRequest(req)
	res, err := s.svr.ListLinkNoJoin(req)
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
	return e.JSON(http.StatusOK, api.ResponsePage(res, hasMore, nextId))
}

func (s *Server) MyAllPostLink(context echo.Context) error {
	return nil
}
