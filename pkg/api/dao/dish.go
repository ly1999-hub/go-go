package dao

import (
	"context"

	"github.com/ly1999-hub/go-go/internal/config/database"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/util/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type Dish struct{}

func (d Dish) InsertOne(ctx context.Context, doc model.Dish) error {
	col := d.getCollection()
	_, err := col.InsertOne(ctx, doc)
	if err != nil {
		log.Error("Error InsertOne-Dao_Dish", log.LogData{"error": err.Error()})
	}
	return nil
}

func (d Dish) InsertMany(ctx context.Context, docs []interface{}) *mongo.InsertManyResult {
	col := d.getCollection()
	results, err := col.InsertMany(ctx, docs)
	if err != nil {
		log.Error("Error InsertMany Dao_Dish", log.LogData{
			"error": err.Error(),
			"docs":  docs,
		})
	}
	return results
}

func (d Dish) getCollection() *mongo.Collection {
	db := database.GetInstance()
	return db.Collection("dish")
}
