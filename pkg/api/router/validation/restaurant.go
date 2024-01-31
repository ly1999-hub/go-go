package validation

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Restaurant struct{}

func (v Restaurant) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload model.RestaurantCreate
		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, err.Error())
		}
		c.Set("restaurant_create", payload)
		return next(c)
	}
}

func (v Restaurant) AllByCustomer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload model.All
		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, err.Error())
		}
		c.Set("restaurant_all", payload)
		return next(c)
	}
}

func (v Restaurant) Detail(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload model.RestaurantDetail
		id := c.Param("id")
		obj, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return response.R400(c, nil, "")
		}
		payload.ID = obj
		c.Set("restaurant_detail", payload)
		return next(c)
	}
}
