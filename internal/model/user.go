package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/ly1999-hub/go-go/internal/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserName  string             `bson:"user_name"`
	Phone     string             `bson:"phone"`
	Email     string             `bson:"email"`
	Login     bool               `bson:"login"`
	Role      string             `bson:"role"`
	Active    bool               `bson:"active"`
	Password  string             `bson:"password"`
	CreatedAt string             `bson:"created_at"`
}

type ResLoginUser struct {
	ID    string
	Token string
}

type UserLoginByPhone struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type ResUserLogin struct {
	ID    string
	Token string
}

type UserCreate struct {
	UserName string `json:"user_name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u UserCreate) Create() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.UserName, validation.Required.Error("không được bỏ trống tên")),
		validation.Field(&u.Password, validation.Required.Error("không được bỏ trống mật khẩu")),
	)
}

func (u UserCreate) ToUser() User {
	doc := User{
		ID:        primitive.NewObjectID(),
		UserName:  u.UserName,
		Phone:     u.Phone,
		Email:     u.Email,
		Password:  util.HashedPassword(u.Password),
		Login:     true,
		Active:    true,
		CreatedAt: time.Now().String(),
	}
	return doc
}

func (u User) ToUserRespone() User {
	return User{
		ID:        u.ID,
		UserName:  u.UserName,
		Phone:     u.Phone,
		Email:     u.Email,
		Login:     u.Login,
		Active:    u.Active,
		CreatedAt: u.CreatedAt,
	}
}
