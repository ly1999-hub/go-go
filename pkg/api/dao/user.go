package dao

import (
	"context"

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

func (d User) getCollection() *mongo.Collection {
	db := database.GetInstance()
	return db.Collection("user")
}
