package initialize

import (
	"os"

	"github.com/ly1999-hub/go-go/internal/config/database"
)

func InitMongo() {
	err := database.Connect(
		os.Getenv("MONGO_HOST"),
		"",
		"",
		os.Getenv("MONGO_DBName"),
		"",
		"",
	)

	if err != nil {
		panic(err)
	}
}
