package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct {
	ID          primitive.ObjectID `bson:"_id"`
	RoleName    string             `bson:"role_name"`
	Description string             `bson:"description"`
	Permissions []string           `bson:"permissions"`
	Code        string             `bson:"code"`
	CreatedAt   string             `bson:"createdAt"`
	UpdatedAt   string             `bson:"updatedAt"`
}

type RoleCreate struct {
	RoleName    string   `json:"role_name"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
	Code        string   `json:"code"`
}

type RoleDelete struct {
	RoleId string `json:"role_id"`
}

type RoleUpdate struct {
	RoleName    string   `json:"role_name"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
	Code        string   `json:"code"`
}

func (r RoleCreate) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.RoleName, validation.Required.Error("khong duoc bo trong name role")),
		validation.Field(&r.Permissions, validation.Required.Error("khong duoc bo trong permission")),
		validation.Field(&r.Code, validation.Required.Error("khong duoc bo trong permission")),
	)
}
