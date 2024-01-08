package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"github.com/ly1999-hub/go-go/internal/util"
	"github.com/ly1999-hub/go-go/pkg/admin/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct{}

func (h Role) Create(c echo.Context) error {
	var (
		s       = service.Role{}
		payload = c.Get("role_create").(model.RoleCreate)
		ctx     = util.GetRequestContext(c)
	)
	res, err := s.Create(ctx, payload)
	if err != nil {
		return response.R400(c, nil, "")
	}
	return response.R200(c, res, "")
}

func (h Role) Delete(c echo.Context) error {
	var (
		s       = service.Role{}
		payload = c.Get("role_delete").(model.RoleDelete)
		ctx     = util.GetRequestContext(c)
	)
	if err := s.Delete(ctx, payload); err != nil {
		return response.R400(c, nil, "")
	}
	return response.R200(c, nil, "")
}

func (h Role) Update(c echo.Context) error {
	var (
		s       = service.Role{}
		payload = c.Get("role_update").(model.RoleUpdate)
		roleId  = c.Get("RoleID").(primitive.ObjectID)
		ctx     = util.GetRequestContext(c)
	)
	err := s.Update(ctx, roleId, payload)
	if err != nil {
		return response.R400(c, nil, "")
	}
	return response.R200(c, nil, "")
}
