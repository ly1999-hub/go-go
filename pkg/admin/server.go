package admin

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ly1999-hub/go-go/internal/middleware/config"
	"github.com/ly1999-hub/go-go/internal/util/log"
	"github.com/ly1999-hub/go-go/pkg/admin/initialize"
	"github.com/ly1999-hub/go-go/pkg/admin/router"
)

func Server(e *echo.Echo) {
	initialize.InitMongo()

	e.Use(config.CORSConfig())
	e.Use(config.RateLimiterConfig())
	e.Use(config.LoggerWithConfig())
	e.Use(middleware.Recover())

	log.Init("server:___")
	router.Init(e)
}
