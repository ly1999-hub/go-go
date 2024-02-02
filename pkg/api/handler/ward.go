package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"github.com/ly1999-hub/go-go/internal/util"
	"github.com/ly1999-hub/go-go/pkg/api/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ward struct{}

func (h Ward) CreateMany(c echo.Context) error {
	var (
		ctx     = util.GetRequestContext(c)
		s       = service.Ward{}
		payload = c.Get("wards_create").(model.WardCreateMany)
	)
	obj, _ := primitive.ObjectIDFromHex(payload.IdDistrict)
	res := s.CreateMany(ctx, payload.Wards, obj)
	if res == -1 {
		return response.R400(c, nil, "")
	}
	return response.R200(c, res, "")
}
