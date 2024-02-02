package validation

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
)

type Province struct{}

func (p Province) CreateMany(netx echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payloads model.ProvinceCreateMany
		if err := c.Bind(&payloads); err != nil {
			return response.R400(c, nil, err.Error())
		}
		c.Set("provinces_create", payloads)
		return netx(c)
	}
}
