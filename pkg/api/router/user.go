package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/pkg/api/handler"
	"github.com/ly1999-hub/go-go/pkg/api/router/middleware"
	"github.com/ly1999-hub/go-go/pkg/api/router/validation"
)

func user(e *echo.Echo) {
	g := e.Group("/user")
	h := handler.User{}
	v := validation.User{}
	vOrder := validation.Order{}
	g.POST("/", h.Create, v.Create)
	g.POST("/login-by-phone", h.LoginByPhone, v.LoginByPhone)
	g.POST("/order", h.Order, middleware.RequireUserLogin, middleware.CheckUserPermission("USER_ORDER"), vOrder.Create)
}
