package database

import (
	"context"
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

// Connect to mongo server
func Connect() error {
	opts := options.Client().ApplyURI(os.Getenv("MONGODB_SRV"))
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		aurora.Red(err.Error())
		fmt.Println(err.Error())
		panic(err)
	}
	db = client.Database(os.Getenv("DATABASE_NAME"))
	return nil
}

// GetInstance ...
func GetInstance() *mongo.Database {
	return db
}
