package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/pkg/api/handler"
	"github.com/ly1999-hub/go-go/pkg/api/router/validation"
)

func customer(e *echo.Echo) {
	g := e.Group("/customer")
	h := handler.Customer{}
	v := validation.Customer{}
	g.POST("/", h.Create, v.Create)
	g.POST("/login-email", h.LoginByEmail, v.LoginByEmail)
	
}
