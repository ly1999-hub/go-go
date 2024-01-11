package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/pkg/admin"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	e := echo.New()
	PORT := os.Getenv("PORT")
	admin.Server(e)

	// log infor start
	e.Logger.Fatal(e.Start(PORT))
}
