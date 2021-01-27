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

func BuildToken(username string, id int) (token string, claims *JwtCustomClaims, err error) {

	// Set custom claims
	claims = &JwtCustomClaims{
		Name: username,
		Id:   id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstr, err := t.SignedString([]byte(secret))

	return tokenstr, claims, err
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
