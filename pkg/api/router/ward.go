package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/pkg/api/handler"
	"github.com/ly1999-hub/go-go/pkg/api/router/validation"
)

func ward(e *echo.Echo) {
	g := e.Group("/ward")
	h := handler.Ward{}
	v := validation.Ward{}

	g.GET("/create-many", h.CreateMany, v.CreateMany)
}
