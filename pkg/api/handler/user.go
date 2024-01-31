package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"github.com/ly1999-hub/go-go/internal/util"
	"github.com/ly1999-hub/go-go/pkg/api/service"
)

type User struct{}

func (u User) Create(c echo.Context) error {
	var (
		s       = service.User{}
		ctx     = util.GetRequestContext(c)
		payload = c.Get("user_create").(model.UserCreate)
	)

	res, err := s.Create(ctx, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, res, "")
}
