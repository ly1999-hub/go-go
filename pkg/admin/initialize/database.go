package initialize

import (
	"github.com/ly1999-hub/go-go/internal/config/database"
)

func InitMongo() {
	err := database.Connect()

	if err != nil {
		panic(err)
	}
}
