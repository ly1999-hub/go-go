package service

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"github.com/ly1999-hub/go-go/internal/util"
	"github.com/ly1999-hub/go-go/internal/util/log"
	daoRole "github.com/ly1999-hub/go-go/pkg/admin/dao"
	"github.com/ly1999-hub/go-go/pkg/api/dao"
	"gopkg.in/mgo.v2/bson"
)

type User struct{}

func (u User) Create(ctx context.Context, payload model.UserCreate) (res *model.ResponseCreate, err error) {
	var (
		dao  = dao.User{}
		role = daoRole.Role{}
	)

	doc := payload.ToUser()
	roleUser := role.FindOne(ctx, bson.M{
		"code": "USER",
	})
	if roleUser.ID.IsZero() {
		return nil, errors.New("không thể xác định quyền!")
	}
	doc.Role = roleUser.Code
	err = dao.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}
	return &model.ResponseCreate{ID: doc.ID.Hex()}, nil
}

func (u User) LoginByPhone(ctx context.Context, payload model.UserLoginByPhone) (*model.ResUserLogin, error) {
	var (
		dao = dao.User{}
	)
	doc := dao.FindOne(ctx, bson.M{"phone": payload.Phone})
	if doc.ID.IsZero() {
		return nil, errors.New(response.CommonNotFound)
	}
	if !doc.Active {
		return nil, errors.New(response.CommonNotActive)
	}
	if !util.CheckPassword(payload.Password, doc.Password) {
		return nil, errors.New(response.CommonErrorPassword)
	}
	return u.generateToken(doc)
}

func (u User) Order(ctx context.Context, payload model.OrderRequest) (*model.ResponseCreate, error) {
	var (
		daoOrder = dao.Order{}
	)
	order := payload.ToOrder()
	order.DishOrder, order.Price = toDishOrder(ctx, payload)
	if err := daoOrder.InsertOne(ctx, order); err != nil {
		return nil, errors.New(response.CommonErrorService)
	}
	return &model.ResponseCreate{ID: order.ID.Hex()}, nil
}

func toDishOrder(ctx context.Context, payload model.OrderRequest) (docs []model.DishOrder, totalPrice int64) {
	var (
		daoDish = dao.Dish{}
	)
	for _, dishOrderRequest := range payload.DishOrderRequest {
		doc := daoDish.FindOne(ctx, bson.M{"_id": util.ObjectIDFromHex(dishOrderRequest.DishOrderId)})
		docs = append(docs, model.DishOrder{
			Dish:   doc,
			Number: dishOrderRequest.Number,
		})
		totalPrice += doc.Price * int64(dishOrderRequest.Number)
	}
	return
}

func (u User) generateToken(user model.User) (*model.ResUserLogin, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id": user.ID,
		"exp": time.Now().Local().Add(time.Second * 15552000).Unix(), // 6 months
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		log.Error("ServiceAdmin-generateToken ", log.LogData{"error": err.Error()})
		return nil, errors.New(response.CommonNotFound)
	}

	return &model.ResUserLogin{ID: user.ID.Hex(), Token: tokenString}, nil
}
