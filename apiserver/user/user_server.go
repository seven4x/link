package user

import (
	"github.com/labstack/echo/v4"
	"github.com/seven4x/link/web/util"

	"net/http"
)

func mvpUser(e echo.Context) error {
	res := make([]UserVO, 0)
	e.JSON(http.StatusOK, util.Success(res))

	return nil
}
