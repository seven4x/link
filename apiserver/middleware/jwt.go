package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt"
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

/*Anonymous
允许匿名访问的接口也需要感知到登陆人
*/
func Anonymous() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, _ := c.Cookie("token")
			if cookie == nil {
				return next(c)
			}
			claims := &JwtCustomClaims{}
			token, err := jwt.ParseWithClaims(cookie.Value, claims, keyFunc)
			if err == nil && token != nil && token.Valid {
				c.Set("user", token)
				return next(c)
			}

			return next(c)
		}
	}
}
func keyFunc(t *jwt.Token) (interface{}, error) {
	// Check the signing method
	if t.Method.Alg() != "HS256" {
		return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
	}

	return []byte(secret), nil
}
