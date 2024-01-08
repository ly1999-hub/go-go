package util

import (
	"context"

	"github.com/labstack/echo/v4"
)

func GetRequestContext(c echo.Context) context.Context {
	return c.Request().Context()
}
