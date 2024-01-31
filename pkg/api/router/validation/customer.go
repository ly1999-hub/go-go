package validation

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
)

type Customer struct{}

func (v Customer) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var customer model.CustomerCreate
		if err := c.Bind(&customer); err != nil {
			return response.R400(c, nil, err.Error())
		}
		c.Set("customer_create", customer)
		return next(c)
	}
}

func (v Customer) LoginByEmail(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var customer model.CustomerLogin
		if err := c.Bind(&customer); err != nil {
			return response.R400(c, nil, err.Error())
		}
		c.Set("customer_login", customer)
		return next(c)
	}
}
