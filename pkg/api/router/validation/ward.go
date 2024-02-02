package validation

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
)

type Ward struct{}

func (v Ward) CreateMany(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload model.WardCreateMany
		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, "")
		}
		c.Set("wards_create", payload)
		return next(c)
	}
}
