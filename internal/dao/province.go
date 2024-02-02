package dao

import (
	"context"

	"github.com/ly1999-hub/go-go/internal/config/database"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/util/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Province struct{}

func (d Province) InsertOne(ctx context.Context, doc model.Province) error {
	col := d.getCollection()
	_, err := col.InsertOne(ctx, doc)
	if err != nil {
		log.Error("Error InsertOne Province_Dao", log.LogData{
			"doc":   doc,
			"Error": err.Error(),
		})
		return err
	}
	return nil
}

func (d Province) InsertMany(ctx context.Context, docs []interface{}, opts ...*options.InsertManyOptions) int {
	col := d.getCollection()
	res, err := col.InsertMany(ctx, docs, opts...)
	if err != nil {
		log.Error("Error InsertMany Province_Dao", log.LogData{
			"docs":  docs,
			"opts":  opts,
			"error": err.Error(),
		})
		return -1
	}
	return len(res.InsertedIDs)
}

func (d Province) getCollection() *mongo.Collection {
	db := database.GetInstance()
	return db.Collection("province")
}
