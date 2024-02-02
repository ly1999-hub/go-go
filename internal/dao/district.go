package dao

import (
	"context"

	"github.com/ly1999-hub/go-go/internal/config/database"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/util/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type District struct{}

func (d District) InsetOne(ctx context.Context, doc model.District) error {
	col := d.getCollection()
	_, err := col.InsertOne(ctx, doc)
	if err != nil {
		log.Error("Error District InsertOne_Dao", log.LogData{
			"doc":   doc,
			"error": err.Error(),
		})
		return err
	}
	return nil
}

func (d District) InsertMany(ctx context.Context, docs []interface{}, opts ...*options.InsertManyOptions) int {
	col := d.getCollection()
	res, err := col.InsertMany(ctx, docs, opts...)
	if err != nil {
		log.Error("Error InsertMany District_Dao", log.LogData{
			"docs":  docs,
			"opts":  opts,
			"error": err.Error(),
		})
		return -1
	}
	return len(res.InsertedIDs)
}

func (d District) getCollection() *mongo.Collection {
	db := database.GetInstance()
	return db.Collection("district")
}
