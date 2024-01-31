package router

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/middleware/jwt"
)

func Init(e *echo.Echo) {
	e.Use(jwt.JWT(os.Getenv("SECRET_KEY")))

	user(e)
	dish(e)
	customer(e)
	restaurant(e)

}
