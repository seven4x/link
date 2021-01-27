package user

import (
	"github.com/Seven4X/link/web/lib/api"
	"github.com/labstack/echo/v4"
	"net/http"
)

func mvpUser(e echo.Context) error {
	res := make([]UserVO, 0)
	e.JSON(http.StatusOK, api.Success(res))

	return nil
}
