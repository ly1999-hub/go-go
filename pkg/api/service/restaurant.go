package service

import (
	"context"
	"errors"
	"sync"

	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"github.com/ly1999-hub/go-go/pkg/api/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Restaurant struct{}

func (s Restaurant) Create(ctx context.Context, payload model.RestaurantCreate, customer model.Customer) (*model.ResponseCreate, error) {
	var (
		daoRestaurant = dao.Restaurant{}
		daoCustomer   = dao.Customer{}
	)
	doc := payload.ToRestaurant()
	doc.Owner = customer.ID
	if err := daoRestaurant.InsertOne(ctx, doc); err != nil {
		return nil, err
	}
	retaurants := append(customer.Restaurants, doc.ID)

	daoCustomer.UpdateByID(ctx, customer.ID, bson.M{"restaurants": retaurants})
	return &model.ResponseCreate{ID: doc.ID.Hex()}, nil
}

func (s Restaurant) AllByCustomer(ctx context.Context, payload model.All, customer model.Customer) (res model.ResponseList) {
	var (
		dao = dao.Restaurant{}
		wg  sync.WaitGroup
	)
	opts := options.Find()
	res.Limit = payload.Limit
	if payload.Limit < 1 {
		res.Limit = 10
	}
	opts.SetLimit(payload.Limit).SetSkip(payload.Page * res.Limit)
	wg.Add(1)
	go func() {
		defer wg.Done()
		count := dao.CountByCondition(ctx, bson.M{})
		res.Total = count
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		docs := dao.FindByCond(ctx, bson.M{"owner": customer.ID}, opts)
		res.List = docs
	}()
	wg.Wait()
	return
}

func (s Restaurant) Detail(ctx context.Context, id primitive.ObjectID, customer model.Customer) (doc model.Restaurant, err error) {
	var (
		dao = dao.Restaurant{}
	)
	doc = dao.FindById(ctx, id)
	if doc.Owner != customer.ID {
		return doc, errors.New(response.CommonUnauthorized)
	}
	if doc.ID.IsZero() {
		return doc, errors.New(response.CommonNotFound)
	}
	return doc, nil
}
