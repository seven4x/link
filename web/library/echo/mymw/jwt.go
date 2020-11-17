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
	Name string `json:"name"`
	Id   int    `json:"id"`
	jwt.StandardClaims
}

func BuildToken(username string, id int) (tokenstr string, claims *JwtCustomClaims) {
	// Set custom claims
	claims = &JwtCustomClaims{
		username,
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	tokenstr, _ = token.SignedString([]byte(secret))

	return tokenstr, claims
}

func JWT() echo.MiddlewareFunc {
	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(secret),
	}

	return middleware.JWTWithConfig(config)
}