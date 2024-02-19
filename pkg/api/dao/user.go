package dao

import (
	"context"
	"fmt"

	"github.com/ly1999-hub/go-go/internal/config/database"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/util/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct{}

func (d User) InsertOne(ctx context.Context, doc model.User) error {
	var col = d.getCollection()
	_, err := col.InsertOne(ctx, doc)
	if err != nil {
		log.Error("Error InsertOne-User", log.LogData{
			"error": err.Error(),
		})
		return err
	}
	return nil
}

func (d User) FindOne(ctx context.Context, filter interface{}) (doc model.User) {
	var col = d.getCollection()

	err := col.FindOne(ctx, filter).Decode(&doc)
	if err != nil {
		log.Error("Error FindOne User-Dao", log.LogData{
			"filter": filter,
			"Error":  err.Error(),
		})
	}
	fmt.Print(doc)
	return
}

func (d User) getCollection() *mongo.Collection {
	db := database.GetInstance()
	return db.Collection("user")
}
