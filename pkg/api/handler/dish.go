package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"github.com/ly1999-hub/go-go/internal/util"
	"github.com/ly1999-hub/go-go/pkg/api/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dish struct{}

func (h Dish) Create(c echo.Context) error {
	var (
		ctx          = util.GetRequestContext(c)
		s            = service.Dish{}
		idRestaurant = c.Param("restaurant")
		payload      = c.Get("dish_create").(model.DishCreate)
		customer     = c.Get("customer").(model.Customer)
	)
	res, err := s.Create(ctx, payload, idRestaurant, customer)
	if err != nil {
		return response.R400(c, res, err.Error())
	}
	return response.R200(c, res, "")
}

func (h Dish) GetAllByRestaurant(c echo.Context) error {
	var (
		ctx          = util.GetRequestContext(c)
		s            = service.Dish{}
		idRestaurant = c.Param("restaurant")
	)
	obj, err := primitive.ObjectIDFromHex(idRestaurant)
	if err != nil {
		return response.R400(c, nil, "")
	}
	res := s.GetAllByRestaurant(ctx, obj)
	return response.R200(c, res, "")
}
