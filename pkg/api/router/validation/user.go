package validation

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
)

type User struct{}

func (u User) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var doc model.UserCreate
		if err := c.Bind(&doc); err != nil {
			return response.R400(c, nil, err.Error())
		}
		if err := doc.Create(); err != nil {
			return response.R400(c, nil, err.Error())
		}
		c.Set("user_create", doc)
		return next(c)
	}
}
