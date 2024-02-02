package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/pkg/api/handler"
	"github.com/ly1999-hub/go-go/pkg/api/router/validation"
)

func province(e *echo.Echo) {
	g := e.Group("/province")

	h := handler.Province{}
	v := validation.Province{}

	g.POST("/create-many", h.CreateMany, v.CreateMany)
}
