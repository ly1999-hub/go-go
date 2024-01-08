package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	ID             primitive.ObjectID `bson:"_id"`
	Name           string             `bson:"name"`
	Email          string             `bson:"email"`
	Phone          string             `bson:"phone"`
	HashedPassword string             `bson:"hashed_password"`
	Birthday       string             `bson:"birthday"`
	Avatar         string             `bson:"avatar"`
	Address        string             `bson:"address"`
	Role           string             `bson:"role,omitempty"`
	Root           bool               `bson:"root"`
	Active         bool               `bson:"active"`
	Login          bool               `bson:"login"`
	CreatedAt      string             `bson:"created_at"`
	UpdatedAt      string             `bson:"updated_at"`
}

type AdminResponse struct {
	Name      string `bson:"name"`
	Email     string `bson:"email"`
	Phone     string `bson:"phone"`
	Birthday  string `bson:"birthday"`
	Avatar    string `bson:"avatar"`
	Address   string `bson:"address"`
	Role      string `bson:"role,omitempty"`
	Active    bool   `bson:"active"`
	Login     bool   `bson:"login"`
	CreatedAt string `bson:"created_at"`
	UpdatedAt string `bson:"updated_at"`
}

type AdminCreate struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Birthday string `json:"birthday"`
	Avatar   string `json:"avatar"`
	Address  string `json:"address"`
	Role     string `json:"role"`
}

type AdminDetail struct {
	ID primitive.ObjectID
}

type AdminForGetPassword struct {
	Email string `json:"email"`
}

// FileUploadInfo ...
type FileUploadInfo struct {
	Filename string `json:"filename"`
	Path     string `json:"path"`
	Ext      string `json:"ext"`
}

func (a AdminCreate) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required.Error("khong duoc bo trong tên")),
		validation.Field(&a.Email, validation.Required.Error("khong duoc bo trong email"),
			is.EmailFormat.Error("dinh dang email")),
		validation.Field(&a.Phone, validation.Required.Error("khong duoc bo trong mật khẩu")),
		validation.Field(&a.Password, validation.Required.Error("khong duoc bo trong mật khẩu")),
		validation.Field(&a.Birthday, validation.Required.Error("khong duoc bo trong ngày sinh")),
		validation.Field(&a.Address, validation.Required.Error("khong duoc bo trong địa chỉ")),
		validation.Field(&a.Role, validation.Required.Error("khong duoc bo trong quyền")),
	)
}

type LoginByEmail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResLoginAdmin struct {
	ID    string `json:"_id"`
	Token string `json:"token"`
}
