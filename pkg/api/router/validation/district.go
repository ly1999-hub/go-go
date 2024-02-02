package validation

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
)

type District struct{}

func (v District) CreateMany(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		var payload model.DistrictCreateMany
		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, err.Error())
		}
		c.Set("districts_create", payload)
		return next(c)
	}
}
