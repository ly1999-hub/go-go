package service

import (
	"context"
	"time"

	"github.com/ly1999-hub/go-go/internal/dao"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type District struct {
}

func (s District) CreateMany(ctx context.Context, payloads []model.District, idProvince primitive.ObjectID) int {
	var (
		daoDistrict = dao.District{}
		daoLog      = dao.Log{}
	)
	docs := make([]interface{}, 0)
	for _, payload := range payloads {
		payload.ID = primitive.NewObjectID()
		payload.Province = idProvince
		docs = append(docs, payload)
	}
	count := daoDistrict.InsertMany(ctx, docs)
	if count == -1 {
		daoLog.InsetOne(ctx, model.Log{
			Type:    response.ERROR,
			Local:   "InsertOne-District",
			Content: docs,
			From:    "",
			Time:    time.Now().String(),
		})
	}
	return count
}
