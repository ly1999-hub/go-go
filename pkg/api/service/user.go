package service

import (
	"context"
	"errors"

	"github.com/ly1999-hub/go-go/internal/model"
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
