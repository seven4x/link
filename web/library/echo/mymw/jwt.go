package mymw

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
)

const (
	secret = "55f0bcd1f387881682359e41"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Token string `json:"token"`
	jwt.StandardClaims
}

func BuildToken(username string, id int) (claims *JwtCustomClaims, err error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	tokenstr, err := token.SignedString([]byte(secret))
	// Set custom claims
	claims = &JwtCustomClaims{
		Name:  username,
		Id:    id,
		Token: tokenstr,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	return claims, err
}

func JWT() echo.MiddlewareFunc {
	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:      &JwtCustomClaims{},
		SigningKey:  []byte(secret),
		TokenLookup: "cookie:token",
	}

	return middleware.JWTWithConfig(config)
}
