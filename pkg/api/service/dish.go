package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/pkg/api/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type Dish struct{}

func (s Dish) Create(ctx context.Context,
	payloads model.DishCreate,
	idRestaurant string,
	customer model.Customer) (*model.ResponseUpdate, error) {
	var (
		daoRestaurant = dao.Restaurant{}
		daoDish       = dao.Dish{}
	)
	obj, err := primitive.ObjectIDFromHex(idRestaurant)
	if err != nil {

		return nil, err
	}
	restaurant := daoRestaurant.FindById(ctx, obj)
	docs := make([]interface{}, 0)
	list := payloads.ToDish()
	for _, doc := range list {
		doc.Restaurant = restaurant.ID
		docs = append(docs, doc)
	}
	res := daoDish.InsertMany(ctx, docs)
	if len(res.InsertedIDs) > 0 {
		for _, obj := range res.InsertedIDs {
			id, ok := obj.(primitive.ObjectID)
			if ok {
				restaurant.Dishs = append(restaurant.Dishs, id)
			}
		}
	}
	if err := daoRestaurant.UpdateById(ctx, restaurant.ID, bson.M{"$set": bson.M{"dishs": restaurant.Dishs, "updated_at": time.Now().String()}}); err != nil {
		return nil, err
	}
	return &model.ResponseUpdate{ID: restaurant.ID.Hex()}, nil
}

func (s Dish) GetAllByRestaurant(ctx context.Context, id primitive.ObjectID) (res model.ResponseList) {
	var (
		dao = dao.Dish{}
		wg  sync.WaitGroup
	)
	fmt.Println(id)
	wg.Add(1)
	go func() {
		defer wg.Done()
		//opts.SetSort(bson.M{"$sort": bson.M{"price": -1}})
		docs := dao.FindByCond(ctx, bson.M{"restaurant": id})
		res.List = docs
		fmt.Println(res.List)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		count := dao.CountByCond(ctx, bson.M{"restaurant": id})
		if count != -1 {
			res.Total = count
			fmt.Println(res.Total)
		}
	}()
	wg.Wait()
	fmt.Println(res)
	return
}
