package setup

import (
	"github.com/Seven4X/link/web/library/echo/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewEcho() (ret *echo.Echo) {
	// Echo instance
	e := echo.New()
	e.Validator = validator.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	return e
}
