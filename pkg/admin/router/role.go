package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/pkg/admin/handler"
	"github.com/ly1999-hub/go-go/pkg/admin/router/middleware"
	"github.com/ly1999-hub/go-go/pkg/admin/router/validation"
)

func role(e *echo.Echo) {
	g := e.Group("/role", middleware.RequireLogin, middleware.CheckPermission("ADMIN"))
	h := handler.Role{}
	v := validation.Role{}

	g.POST("/", h.Create, v.Create)
	g.DELETE("/", h.Delete, v.Delete)
	g.PATCH("/:id", h.Update, v.Update)
}
