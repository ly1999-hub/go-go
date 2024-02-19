package dao

import (
	"context"

	"github.com/ly1999-hub/go-go/internal/config/database"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/util/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type Order struct{}

func (d Order) InsertOne(ctx context.Context, doc model.Order) error {
	col := d.getCollection()
	_, err := col.InsertOne(ctx, doc)
	if err != nil {
		log.Error("Error InsertOne-Order", log.LogData{
			"error": err.Error(),
		})
		return err
	}
	return nil
}

func (d Order) getCollection() *mongo.Collection {
	db := database.GetInstance()
	return db.Collection("order")
}
