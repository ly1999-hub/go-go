package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"github.com/ly1999-hub/go-go/internal/util/log"
	"github.com/ly1999-hub/go-go/pkg/admin/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type Role struct{}

func (s Role) Create(ctx context.Context, payload model.RoleCreate) (*model.ResponseCreate, error) {
	var (
		daoRole = dao.Role{}
	)
	role := daoRole.FindOne(ctx, bson.M{"code": payload.Code})
	fmt.Print("role", role)
	if !role.ID.IsZero() {
		return nil, errors.New(response.CommonExistEmail)
	}
	doc := model.Role{
		ID:          primitive.NewObjectID(),
		RoleName:    payload.RoleName,
		Description: payload.Description,
		Permissions: payload.Permissions,
		Code:        payload.Code,
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}
	fmt.Print(doc)
	if err := daoRole.InsertOne(ctx, doc); err != nil {
		return nil, errors.New(response.CommonBadRequest)
	}
	return &model.ResponseCreate{ID: doc.ID.Hex()}, nil
}

func (s Role) Delete(ctx context.Context, payload model.RoleDelete) error {
	var (
		dao = dao.Role{}
	)
	id, err := primitive.ObjectIDFromHex(payload.RoleId)
	if err != nil {
		log.Error("serviceRole-Create", log.LogData{
			"error": err.Error(),
		})
		return err
	}
	if err := dao.DeleteByID(ctx, id); err != nil {
		log.Error("serviceRole-Update", log.LogData{
			"error": err.Error(),
		})
		return err
	}
	return nil
}

func (s Role) Update(ctx context.Context, roleID primitive.ObjectID, payload model.RoleUpdate) error {
	var (
		dao = dao.Role{}
	)

	doc := bson.M{
		"$set": bson.M{
			"role_name":   payload.RoleName,
			"description": payload.Description,
			"permissions": payload.Permissions,
			"code":        payload.Code,
			"updated_at":  time.Now().String(),
		},
	}
	err := dao.UpdateByID(ctx, roleID, doc)
	if err != nil {
		log.Error("serviceRole-Update", log.LogData{
			"error": err.Error(),
		})
		return err
	}
	return nil
}

func (s Role) All(c context.Context, page model.All) (res model.ResponseList) {
	var (
		dao = dao.Role{}
	)
	result := dao.FindAll(c, bson.M{})
	res.List = result
	return
}
