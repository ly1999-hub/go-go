package dao

import (
	"context"

	"github.com/ly1999-hub/go-go/internal/config/database"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/util/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (d Dish) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (doc model.Dish) {
	col := d.getCollection()
	if err := col.FindOne(ctx, filter, opts...).Decode(&doc); err != nil {
		log.Error("Error FindOne Dish_Dao", log.LogData{
			"filter": filter,
			"opts":   opts,
			"Error":  err.Error(),
		})

	}
	return
}

func (d Dish) FindByCond(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (docs []model.Dish) {
	col := d.getCollection()
	cursor, err := col.Find(ctx, filter, opts...)
	if err != nil {
		log.Error("Error FindByCond Dish_Dao Find", log.LogData{
			"filter": filter,
			"opts":   opts,
			"error":  err.Error(),
		})
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &docs); err != nil {
		log.Error("Error FindByCond Dish_Dao All", log.LogData{
			"error": err.Error(),
		})
	}
	return
}

func (d Dish) CountByCond(ctx context.Context, filter interface{}, opts ...*options.CountOptions) int64 {
	col := d.getCollection()
	cont, err := col.CountDocuments(ctx, filter, opts...)
	if err != nil {
		log.Error("Error CountByCond Dish_Dao", log.LogData{
			"filter": filter,
			"opts":   opts,
			"error":  err.Error(),
		})
		return -1
	}
	return cont
}

func (d Dish) UpdatedOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) error {
	col := d.getCollection()
	_, err := col.UpdateOne(ctx, filter, update, opts...)
	if err != nil {
		log.Error("ERROR UpdateOne Dish_Dao", log.LogData{
			"filter": filter,
			"update": update,
			"opts":   opts,
			"Error":  err.Error(),
		})
		return err
	}
	return nil
}

func (d Dish) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) error {
	col := d.getCollection()
	_, err := col.DeleteOne(ctx, filter, opts...)
	if err != nil {
		log.Error("ERROR DeleteOne Dish_Dao", log.LogData{
			"filter": filter,
			"opts":   opts,
			"error":  err.Error(),
		})
		return err
	}
	return nil
}

func (d Dish) getCollection() *mongo.Collection {
	db := database.GetInstance()
	return db.Collection("dish")
}
