package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"github.com/ly1999-hub/go-go/internal/util"
	"github.com/ly1999-hub/go-go/pkg/api/service"
)

type Province struct{}

func (h Province) CreateMany(c echo.Context) error {
	var (
		ctx     = util.GetRequestContext(c)
		s       = service.Province{}
		payload = c.Get("provinces_create").(model.ProvinceCreateMany)
	)
	res := s.CreateMany(ctx, payload.Provinces)
	fmt.Println(payload.Provinces)
	if res == -1 {
		return response.R400(c, res, "")
	}
	return response.R200(c, res, "")
}
