package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/pkg/admin/handler"
	"github.com/ly1999-hub/go-go/pkg/admin/middleware"
	"github.com/ly1999-hub/go-go/pkg/admin/validation"
)

func admin(e *echo.Echo) {
	g := e.Group("/admin")
	h := handler.Admin{}
	v := validation.Admin{}

	g.GET("/", h.GetMe, middleware.RequireLogin)
	g.GET("/:id", h.GetDetail, middleware.RequireLogin, middleware.CheckPermission("ADMIN_VIEW"), v.GetDetail)
	g.GET("/all", h.All, v.All)

	g.POST("/", h.Create, middleware.RequireLogin, middleware.CheckPermission("ADMIN_CREATE"), v.Create)
	g.POST("/login-email", h.LoginEmail, v.LoginEmail)
	g.POST("/forget-password", h.ForGetPassword, v.ForGetPassword)
	g.POST("/upload-avatar", h.UploadAvatar, middleware.RequireLogin, middleware.UploadSingleFile)
}
