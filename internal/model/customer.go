package model

import (
	"time"

	"github.com/ly1999-hub/go-go/internal/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID          primitive.ObjectID   `bson:"_id"`
	UserName    string               `bson:"user_name"`
	Phone       string               `bson:"phone"`
	Email       string               `bson:"email"`
	Avatar      string               `bson:"avatar"`
	Login       bool                 `bson:"login"`
	Role        string               `bson:"role"`
	Active      bool                 `bson:"active"`
	Address     string               `bson:"address"`
	Restaurants []primitive.ObjectID `bson:"restaurants"`
	Password    string               `bson:"password"`
	CreatedAt   string               `bson:"created_at"`
	UpdatedAt   string               `bson:"updated_at"`
}

type CustomerCreate struct {
	UserName string `json:"user_name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

type CustomerLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c CustomerCreate) ToCustomer() Customer {
	return Customer{
		ID:        primitive.NewObjectID(),
		UserName:  c.UserName,
		Phone:     c.Phone,
		Email:     c.Email,
		Address:   c.Address,
		Login:     true,
		Active:    true,
		Role:      "",
		Password:  util.HashedPassword(c.Password),
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}

}
