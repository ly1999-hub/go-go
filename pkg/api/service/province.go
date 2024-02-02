package service

import (
	"context"
	"time"

	"github.com/ly1999-hub/go-go/internal/dao"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Province struct{}

func (s Province) CreateMany(ctx context.Context, payloads []model.Province) int {
	var (
		daoProvince = dao.Province{}
		daoLog      = dao.Log{}
	)
	docs := make([]interface{}, 0)
	for _, payload := range payloads {
		payload.ID = primitive.NewObjectID()
		docs = append(docs, payload)
	}
	res := daoProvince.InsertMany(ctx, docs)
	if res == -1 {
		daoLog.InsetOne(ctx, model.Log{
			Type:    response.ERROR,
			Local:   "InsertOne-Province",
			Content: docs,
			From:    "",
			Time:    time.Now().String(),
		})
	}
	return res
}
