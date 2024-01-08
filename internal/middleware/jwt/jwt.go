package jwt

import (
	jwtecho "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func JWT(secretKey string) echo.MiddlewareFunc {
	return jwtecho.WithConfig(jwtecho.Config{
		SigningKey: []byte(secretKey),
		Skipper: func(c echo.Context) bool {
			authHeader := c.Request().Header.Get("Authorization") // get header with authorization
			// return true to skip middleware
			return authHeader == "" || authHeader == "Bearer" || authHeader == "Bearer "
		},
	})
}
