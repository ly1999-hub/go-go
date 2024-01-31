package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/ly1999-hub/go-go/internal/util"
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

type AdminChangePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
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

func (login LoginByEmail) Validate() error {
	return validation.ValidateStruct(&login,
		validation.Field(&login.Email, validation.Required.Error("Không được bỏ trống email."), is.EmailFormat.Error("định dạng là email: *@gmail.com")),
		validation.Field(&login.Password, validation.Length(4, 50).Error("mật khẩu nhiều hơn 4 ký tự và ít hơn 50 ký tự")),
	)
}

type ResLoginAdmin struct {
	ID    string `json:"_id"`
	Token string `json:"token"`
}

func (a AdminCreate) NewAdmin() Admin {
	return Admin{
		ID:             primitive.NewObjectID(),
		Name:           a.Name,
		Email:          a.Email,
		Phone:          a.Phone,
		HashedPassword: util.HashedPassword(a.Password),
		Birthday:       a.Birthday,
		Avatar:         a.Avatar,
		Address:        a.Address,
		Role:           a.Role,
		Root:           false,
		Active:         true,
		CreatedAt:      time.Now().String(),
		UpdatedAt:      time.Now().String(),
	}
}
