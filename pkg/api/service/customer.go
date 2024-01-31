package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"github.com/ly1999-hub/go-go/internal/util"
	"github.com/ly1999-hub/go-go/internal/util/log"
	daoAdmin "github.com/ly1999-hub/go-go/pkg/admin/dao"
	"github.com/ly1999-hub/go-go/pkg/api/dao"
	"gopkg.in/mgo.v2/bson"
)

type Customer struct{}

func (s Customer) Create(ctx context.Context, payload model.CustomerCreate) (*model.ResponseCreate, error) {
	var (
		dao     = dao.Customer{}
		daoRole = daoAdmin.Role{}
	)

	fmt.Println(payload)
	customer := s.CheckExitByEmail(ctx, payload.Email)
	if !customer.ID.IsZero() {
		return nil, errors.New(response.CommonExistEmail)
	}
	doc := payload.ToCustomer()
	role := daoRole.FindOne(ctx, bson.M{"code": "CUSTOMER"})
	if role.ID.IsZero() {
		return nil, errors.New("Not found Role-Customer")
	}
	doc.Role = role.Code
	dao.InsertOne(ctx, doc)
	return &model.ResponseCreate{ID: doc.ID.Hex()}, nil
}

func (s Customer) CheckExitByEmail(ctx context.Context, email string) model.Customer {
	var (
		dao = dao.Customer{}
	)

	doc := dao.FindOne(ctx, bson.M{"email": email})
	return doc
}

func (s Customer) LoginByEmail(ctx context.Context, payload model.CustomerLogin) (*model.ResLoginAdmin, error) {
	customer := s.CheckExitByEmail(ctx, payload.Email)
	if customer.ID.IsZero() {
		return nil, errors.New(response.CommonNotFound)
	}
	if !util.CheckPassword(payload.Password, customer.Password) {
		return nil, errors.New(response.CommonErrorPassword)
	}
	if !customer.Active {
		return nil, errors.New(response.CommonNotActive)
	}
	return generateToken(customer)

}

func generateToken(customer model.Customer) (*model.ResLoginAdmin, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id": customer.ID,
		"exp": time.Now().Local().Add(time.Second * 15552000).Unix(), // 6 months
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	fmt.Print(tokenString)
	if err != nil {
		log.Error("ServiceAdmin-generateToken ", log.LogData{"error": err.Error()})
		return nil, errors.New(response.CommonNotFound)
	}

	return &model.ResLoginAdmin{ID: customer.ID.Hex(), Token: tokenString}, nil
}
