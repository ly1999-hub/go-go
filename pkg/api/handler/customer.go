package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"github.com/ly1999-hub/go-go/internal/util"
	"github.com/ly1999-hub/go-go/pkg/api/service"
)

type Customer struct{}

func (h Customer) Create(c echo.Context) error {
	var (
		ctx     = util.GetRequestContext(c)
		s       = service.Customer{}
		payload = c.Get("customer_create").(model.CustomerCreate)
	)
	res, err := s.Create(ctx, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, res, "")

}

func (h Customer) LoginByEmail(c echo.Context) error {
	var (
		ctx     = util.GetRequestContext(c)
		s       = service.Customer{}
		payload = c.Get("customer_login").(model.CustomerLogin)
	)
	res, err := s.LoginByEmail(ctx, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, res, "")

}
