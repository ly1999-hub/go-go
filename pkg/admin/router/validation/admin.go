package validation

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct{}

func (a Admin) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var adminCreate model.AdminCreate
		if err := c.Bind(&adminCreate); err != nil {
			return response.R400(c, nil, "")
		}
		if err := adminCreate.Validate(); err != nil {
			return response.R400(c, nil, err.Error())
		}
		c.Set("admin_create", adminCreate)
		return next(c)
	}
}

func (a Admin) LoginEmail(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload model.LoginByEmail
		if err := c.Bind(&payload); err != nil {
			return response.R400(c, "", err.Error())
		}
		if err := payload.Validate(); err != nil {
			return response.R400(c, nil, err.Error())
		}
		c.Set("admin_login", payload)
		return next(c)
	}
}

func (a Admin) All(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var page model.All
		if err := c.Bind(&page); err != nil {
			return response.R400(c, nil, "")
		}
		c.Set("all", page)
		fmt.Print(page)
		return next(c)
	}
}

func (a Admin) ForGetPassword(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload model.AdminForGetPassword
		if err := c.Bind(&payload); err != nil {
			return response.R400(c, "", err.Error())
		}
		c.Set("admin_forget_password", payload)
		return next(c)
	}
}

func (a Admin) GetDetail(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id      = c.Param("id")
			payload = model.AdminDetail{}
		)
		idObject, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return response.R400(c, nil, response.CommonErrorService)
		}
		payload.ID = idObject

		c.Set("admin_detail", payload)
		return next(c)
	}
}

func (a Admin) ChangePassword(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload = model.AdminChangePassword{}
		)
		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, err.Error())
		}

		c.Set("admin_change_password", payload)
		return next(c)
	}
}
