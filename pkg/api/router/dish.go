package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/pkg/api/handler"
	"github.com/ly1999-hub/go-go/pkg/api/router/middleware"
	"github.com/ly1999-hub/go-go/pkg/api/router/validation"
)

func dish(e *echo.Echo) {
	g := e.Group("/dish", middleware.RequireLogin, middleware.CheckPermission("RESTAURANT_EDIT"))
	h := handler.Dish{}
	v := validation.Dish{}

	g.POST("/:restaurant/create", h.Create, v.Create)
}
