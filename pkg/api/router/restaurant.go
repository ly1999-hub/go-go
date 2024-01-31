package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/pkg/api/handler"
	"github.com/ly1999-hub/go-go/pkg/api/router/middleware"
	"github.com/ly1999-hub/go-go/pkg/api/router/validation"
)

func restaurant(e *echo.Echo) {
	g := e.Group("/retaurant")
	h := handler.Restaurant{}
	v := validation.Restaurant{}
	g.POST("/", h.Create, middleware.RequireLogin, middleware.CheckPermission("RESTAURANT_CREATE"), v.Create)
	g.GET("/all-by-customer", h.AllByCustomer, middleware.RequireLogin, middleware.CheckPermission("RESTAURANT_VIEW"), v.AllByCustomer)
	g.GET("/:id", h.Detail, middleware.RequireLogin, middleware.CheckPermission("RESTAURANT_VIEW"), v.Detail)
}
