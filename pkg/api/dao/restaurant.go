package dao

import (
	"context"

	"github.com/ly1999-hub/go-go/internal/config/database"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/util/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Restaurant struct{}

func (d Restaurant) InsertOne(ctx context.Context, doc model.Restaurant) error {
	var col = d.getCollection()
	_, err := col.InsertOne(ctx, doc)
	if err != nil {
		log.Error("ERROR InsertOnen Restaurant", log.LogData{"error": err.Error()})
		return err
	}
	return nil
}

func (d Restaurant) FindOne(ctx context.Context, filter interface{}) (doc model.Restaurant) {
	col := d.getCollection()
	if err := col.FindOne(ctx, filter).Decode(&doc); err != nil {
		log.Error("ERROR FINDONE RESTAURANT", log.LogData{"error": err.Error()})
	}
	return doc
}

func (d Restaurant) FindById(ctx context.Context, id primitive.ObjectID) (doc model.Restaurant) {
	return d.FindOne(ctx, bson.M{"_id": id})
}

func (d Restaurant) FindByCond(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (doc []model.Restaurant) {
	var col = d.getCollection()
	cursor, err := col.Find(ctx, filter, opts...)
	if err != nil {
		log.Error("Error FindByCond-Restaurant DAO-Find", log.LogData{"error": err.Error()})
		return
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &doc); err != nil {
		log.Error("Error FindByCond-Restaurant Dao-All", log.LogData{"error": err.Error()})
		return
	}
	return doc
}

func (d Restaurant) CountByCondition(ctx context.Context, cond interface{}) int64 {
	col := d.getCollection()
	total, err := col.CountDocuments(ctx, cond)
	if err != nil {
		log.Error("dao.Province - CountByCondition", log.LogData{
			"err":  err.Error(),
			"cond": cond,
		})
	}
	return total
}

func (d Restaurant) UpdateOne(ctx context.Context, filter interface{}, cond interface{}) error {
	col := d.getCollection()
	_, err := col.UpdateOne(ctx, filter, cond)
	if err != nil {
		log.Error("Error UpdateOne Restaurant_Dao", log.LogData{"Error": err.Error()})
		return err
	}
	return nil
}

func (d Restaurant) UpdateById(ctx context.Context, id primitive.ObjectID, cond interface{}) error {
	return d.UpdateOne(ctx, bson.M{"_id": id}, cond)
}
func (d Restaurant) getCollection() mongo.Collection {
	db := database.GetInstance()
	return *db.Collection("restaurant")
}
