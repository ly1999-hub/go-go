package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"github.com/ly1999-hub/go-go/internal/util"
	"github.com/ly1999-hub/go-go/pkg/api/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type District struct{}

func (d District) CreateMany(c echo.Context) error {
	var (
		ctx     = util.GetRequestContext(c)
		s       = service.District{}
		payload = c.Get("districts_create").(model.DistrictCreateMany)
	)
	obj, _ := primitive.ObjectIDFromHex(payload.IdProvince)
	res := s.CreateMany(ctx, payload.Districts, obj)
	return response.R200(c, res, "")
}
