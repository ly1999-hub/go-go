package dao

import (
	"context"

	"github.com/ly1999-hub/go-go/internal/config/database"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/util/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type Customer struct{}

func (d Customer) InsertOne(ctx context.Context, doc model.Customer) error {
	var col = d.getCollection()
	_, err := col.InsertOne(ctx, doc)

	if err != nil {
		log.Error("Error InsertOne Customer", log.LogData{"error: ": err.Error()})
		return err
	}
	return nil
}

func (d Customer) FindOne(ctx context.Context, cound interface{}) (doc model.Customer) {
	var col = d.getCollection()

	if err := col.FindOne(ctx, cound).Decode(&doc); err != nil {
		log.Error("Error Dao-FindOne Customer", log.LogData{
			"err:": err.Error(),
		})
	}
	return doc
}

func (d Customer) FindByID(ctx context.Context, id primitive.ObjectID) model.Customer {
	return d.FindOne(ctx, bson.M{"_id": id})
}

func (d Customer) UpdateOne(ctx context.Context, cond, payload interface{}) error {
	var col = d.getCollection()
	if _, err := col.UpdateOne(ctx, cond, payload); err != nil {
		log.Error("Error UpdateOne Customer-dao", log.LogData{
			"error": err.Error(),
		})
		return err
	}
	return nil
}

func (d Customer) UpdateByID(ctx context.Context, id primitive.ObjectID, payload interface{}) error {
	return d.UpdateOne(ctx, bson.M{"_id": id}, payload)
}

func (d Customer) getCollection() *mongo.Collection {
	db := database.GetInstance()
	return db.Collection("customer")
}
