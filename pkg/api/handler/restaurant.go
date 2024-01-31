package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"github.com/ly1999-hub/go-go/internal/util"
	"github.com/ly1999-hub/go-go/pkg/api/service"
)

type Restaurant struct{}

func (h Restaurant) Create(c echo.Context) error {
	var (
		ctx      = util.GetRequestContext(c)
		s        = service.Restaurant{}
		payload  = c.Get("restaurant_create").(model.RestaurantCreate)
		customer = c.Get("customer").(model.Customer)
	)
	res, err := s.Create(ctx, payload, customer)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, res, "")
}

func (h Restaurant) AllByCustomer(c echo.Context) error {
	var (
		ctx      = util.GetRequestContext(c)
		s        = service.Restaurant{}
		payload  = c.Get("restaurant_all").(model.All)
		customer = c.Get("customer").(model.Customer)
	)
	res := s.AllByCustomer(ctx, payload, customer)
	return response.R200(c, res, "")
}

func (h Restaurant) Detail(c echo.Context) error {
	var (
		ctx      = util.GetRequestContext(c)
		s        = service.Restaurant{}
		id       = c.Get("restaurant_detail").(model.RestaurantDetail)
		customer = c.Get("customer").(model.Customer)
	)
	res, err := s.Detail(ctx, id.ID, customer)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, res, "")
}
