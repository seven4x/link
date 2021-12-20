package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/seven4x/link/web/middleware"
)

func GetUserId(e echo.Context) int {
	u := e.Get(User)
	if u == nil {
		return 0
	}
	user := u.(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	return claims.Id
}
func GetUser(e echo.Context) *middleware.JwtCustomClaims {
	u := e.Get(User)
	if u == nil {
		return nil
	}
	user := u.(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	return claims
}
