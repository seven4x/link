package app

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetUserId(e echo.Context) int {
	u := e.Get(User)
	if u == nil {
		return 0
	}
	user := u.(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims.Id
}
func GetUser(e echo.Context) *JwtCustomClaims {
	u := e.Get(User)
	if u == nil {
		return nil
	}
	user := u.(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims
}
