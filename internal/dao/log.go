package dao

import (
	"context"

	"github.com/ly1999-hub/go-go/internal/config/database"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/util/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type Log struct{}

func (d Log) InsetOne(ctx context.Context, doc model.Log) error {
	col := d.getCollection()
	_, err := col.InsertOne(ctx, doc)
	if err != nil {
		log.Error("Error InsertOne Log_Dao", log.LogData{
			"doc":   doc,
			"error": err.Error(),
		})
		return nil
	}
	return nil
}

func (d Log) getCollection() *mongo.Collection {
	db := database.GetInstance()
	return db.Collection("logs")
}
