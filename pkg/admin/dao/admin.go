package dao

import (
	"context"
	"fmt"

	"github.com/ly1999-hub/go-go/internal/config/database"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/util/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Admin struct{}

func (d Admin) InsertOne(ctx context.Context, doc model.Admin) error {
	col := d.getCollection()
	_, err := col.InsertOne(ctx, doc)
	if err != nil {
		fmt.Println("error-insertOne Admin")
		log.Error("dao.User - UpdateOne", log.LogData{
			"doc": doc,
			"err": err.Error(),
		})
	}
	return err
}

func (d Admin) FindOne(ctx context.Context, cond interface{}) (doc model.Admin) {
	col := d.getCollection()
	if err := col.FindOne(ctx, cond).Decode(&doc); err != nil {
		fmt.Println(err.Error())
		log.Error("find-one", log.LogData{
			"cond":  cond,
			"error": err.Error(),
		})
	}
	return doc
}

func (d Admin) FindByID(ctx context.Context, id primitive.ObjectID) (doc model.Admin) {
	return d.FindOne(ctx, bson.M{"_id": id})
}

// CountByCondition ...
func (d Admin) CountByCondition(ctx context.Context, cond interface{}) int64 {
	col := d.getCollection()
	total, err := col.CountDocuments(ctx, cond)
	if err != nil {
		fmt.Println(err.Error())
		log.Error("dao.Province - CountByCondition", log.LogData{
			"err":  err.Error(),
			"cond": cond,
		})
	}
	return total
}

func (d Admin) UpdateOne(ctx context.Context, cond, payload interface{}) error {
	var col = d.getCollection()

	_, err := col.UpdateOne(ctx, cond, payload)
	if err != nil {
		fmt.Println(err.Error())
		log.Error("UpdateOne Admin-daoAdmin", log.LogData{
			"cond":    cond,
			"payload": payload,
			"error":   err.Error(),
		})
	}
	return err
}

func (d Admin) UpdateByID(ctx context.Context, id primitive.ObjectID, payload interface{}) error {
	return d.UpdateOne(ctx, bson.M{"_id": id}, payload)
}

func (d Admin) FindByCondtion(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.Admin) {
	col := d.getCollection()
	cursor, err := col.Find(ctx, cond, opts...)
	if err != nil {
		fmt.Println(err.Error())
		log.Error("dao.Admin- FindByCondtion", log.LogData{
			"Cond":  cond,
			"Opts":  opts,
			"Error": err.Error(),
		})
		return
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &docs); err != nil {
		fmt.Println(err.Error())
		log.Error("dao.Province - FindByCondition - decode", log.LogData{
			"cond": cond,
			"opts": opts,
			"err":  err.Error(),
		})
		return
	}
	return docs
}

func (d Admin) DeleteOne(ctx context.Context, filter interface{}) error {
	var col = d.getCollection()
	_, err := col.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (d Admin) DeleteByID(ctx context.Context, id primitive.ObjectID) error {
	return d.DeleteOne(ctx, bson.M{"_id": id})
}

func (d Admin) getCollection() *mongo.Collection {
	db := database.GetInstance()
	return db.Collection("admins")
}
