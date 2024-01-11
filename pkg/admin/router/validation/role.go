package validation

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct{}

func (v Role) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload model.RoleCreate
		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, "")
		}
		c.Set("role_create", payload)
		return next(c)
	}
}

func (v Role) Delete(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload model.RoleDelete
		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, "")
		}
		c.Set("role_delete", payload)
		return next(c)
	}
}

func (v Role) Update(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.RoleUpdate
		)
		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, "")
		}
		roleID := c.Param("id")
		objectID, err := primitive.ObjectIDFromHex(roleID)
		if err != nil {
			return response.R400(c, nil, "")
		}

		c.Set("RoleID", objectID)
		c.Set("role_update", payload)
		fmt.Print(roleID, payload)
		return next(c)
	}
}

func (v Role) Detail(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload model.AdminDetail

		idString := c.Param("id")
		if idString == "" {
			return response.R400(c, nil, "")
		}
		idObject, err := primitive.ObjectIDFromHex(idString)
		if err != nil {
			return response.R400(c, nil, "")
		}
		payload.ID = idObject
		c.Set("admin_detail", payload)
		return next(c)
	}
}
