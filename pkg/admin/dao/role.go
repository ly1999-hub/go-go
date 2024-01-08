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

type Role struct{}

func (d Role) InsertOne(ctx context.Context, doc model.Role) error {
	col := d.getCollection()
	_, err := col.InsertOne(ctx, doc)
	if err != nil {
		log.Error("dao.Role - InsertOne", log.LogData{
			"doc": doc,
			"err": err.Error(),
		})
	}
	return err
}

func (d Role) FindOne(ctx context.Context, cond interface{}) (role model.Role) {
	col := d.getCollection()
	if err := col.FindOne(ctx, cond).Decode(&role); err != nil {
		log.Error("dao.Role - FindOne", log.LogData{
			"cond": cond,
			"err":  err.Error(),
		})
	}
	return role
}

func (d Role) FindByID(ctx context.Context, id primitive.ObjectID) (doc model.Role) {
	return d.FindOne(ctx, bson.M{"_id": id})
}

func (d Role) UpdateByID(ctx context.Context, Id primitive.ObjectID, payload ...interface{}) error {
	var col = d.getCollection()
	_, err := col.UpdateByID(ctx, Id, payload)

	if err != nil {
		log.Error("dao.Role - updateByID", log.LogData{
			"id":      Id,
			"payload": payload,
		})
		return err
	}
	return nil
}

func (d Role) UpdateOne(ctx context.Context, filter, count interface{}) error {
	var col = d.getCollection()

	_, err := col.UpdateOne(ctx, filter, count)

	if err != nil {
		log.Error("daoRole - UpdateOne", log.LogData{
			"filter": filter,
			"count":  count,
		})
		return err
	}
	return nil
}

func (d Role) DeleteByID(ctx context.Context, Id primitive.ObjectID) error {
	var col = d.getCollection()

	_, err := col.DeleteOne(ctx, bson.M{"_id": Id})

	if err != nil {
		log.Error("daoRole - DeleteByID", log.LogData{
			"id": Id,
		})
		return err
	}
	return nil
}

func (d Role) FindAll(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.Role) {
	var col = d.getCollection()
	cursor, err := col.Find(ctx, cond, opts...)
	if err != nil {
		log.Error("daoRole - findAll", log.LogData{
			"cond": cond,
			"opts": opts,
		})
		return
	}

	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &docs); err != nil {
		return
	}
	return

}

func (d Role) getCollection() *mongo.Collection {
	db := database.GetInstance()
	return db.Collection("roles")
}
