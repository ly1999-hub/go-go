package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/pkg/api/handler"
	"github.com/ly1999-hub/go-go/pkg/api/router/validation"
)

func district(e *echo.Echo) {
	g := e.Group("/district")
	h := handler.District{}
	v := validation.District{}

	g.POST("/create-many", h.CreateMany, v.CreateMany)
}
