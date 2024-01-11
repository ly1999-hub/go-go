package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"github.com/ly1999-hub/go-go/internal/util"
	"github.com/ly1999-hub/go-go/pkg/admin/service"
)

type Admin struct{}

func (a Admin) Create(c echo.Context) error {
	var (
		payload = c.Get("admin_create").(model.AdminCreate)
		ctx     = util.GetRequestContext(c)
		s       = service.Admin{}
	)
	res, err := s.Create(ctx, payload)
	if err != nil {
		return response.R400(c, nil, "")
	}
	return response.R200(c, res, "")
}

func (a Admin) GetMe(c echo.Context) error {
	var (
		s   = service.Admin{}
		me  = c.Get("admin").(model.Admin)
		ctx = util.GetRequestContext(c)
	)
	res := s.GetMe(ctx, me)
	return response.R200(c, res, "")
}

func (a Admin) GetDetail(c echo.Context) error {
	var (
		s       = service.Admin{}
		admin   = c.Get("admin").(model.Admin)
		ctx     = util.GetRequestContext(c)
		payload = c.Get("admin_detail").(model.AdminDetail)
	)

	res, err := s.GetDetail(ctx, admin, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, res, "")
}

func (a Admin) LoginEmail(c echo.Context) error {
	var (
		payload = c.Get("admin_login").(model.LoginByEmail)
		ctx     = util.GetRequestContext(c)
		s       = service.Admin{}
	)
	res, err := s.LoginByEmail(ctx, payload)
	if err != nil {
		return response.R400(c, payload, err.Error())
	}
	return response.R200(c, res, "")
}

func (a Admin) All(c echo.Context) error {
	var (
		payload = c.Get("all").(model.All)
		ctx     = util.GetRequestContext(c)
		s       = service.Admin{}
	)

	res := s.GetAll(ctx, payload)
	return response.R200(c, res, "")
}

func (a Admin) ForGetPassword(c echo.Context) error {
	var (
		payload = c.Get("admin_forget_password").(model.AdminForGetPassword)
		ctx     = util.GetRequestContext(c)
		s       = service.Admin{}
	)
	res, err := s.ForGetPassword(ctx, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}

	return response.R200(c, res, "")
}

func (a Admin) UploadAvatar(c echo.Context) error {
	var (
		s          = service.Admin{}
		admin      = c.Get("admin").(model.Admin)
		fileUpload = c.Get("file_avatar").(model.FileUploadInfo)
		ctx        = util.GetRequestContext(c)
	)
	res, err := s.UploadAvatar(ctx, admin, fileUpload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, res, "")
}

func (a Admin) ChangePassword(c echo.Context) error {
	var (
		s       = service.Admin{}
		me      = c.Get("admin").(model.Admin)
		ctx     = util.GetRequestContext(c)
		payload = c.Get("admin_change_password").(model.AdminChangePassword)
	)
	res, err := s.ChangePassword(ctx, me, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, res, "")
}
