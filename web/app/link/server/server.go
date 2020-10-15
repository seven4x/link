package server

import (
	"github.com/Seven4X/link/web/library/config"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Router(e *echo.Echo) {
	g := e.Group("/link")
	g.GET("/preview-token", getPreviewToken)
}

func getPreviewToken(e echo.Context) error {
	str := config.GetString(config.LinkPreviewToken)

	_ = e.HTML(http.StatusOK, str)
	return nil
}
