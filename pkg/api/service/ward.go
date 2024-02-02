package service

import (
	"context"
	"sync"
	"time"

	"github.com/ly1999-hub/go-go/internal/dao"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ward struct{}

func (s Ward) CreateMany(ctx context.Context, payloads []model.Ward, idDistrict primitive.ObjectID) int {
	var (
		daoWard = dao.Ward{}
		daoLog  = dao.Log{}
		wg      sync.WaitGroup
	)
	docs := make([]interface{}, 0)

	for _, payload := range payloads {
		wg.Add(1)
		go func(payload model.Ward) {
			defer wg.Done()
			payload.District = idDistrict
			docs = append(docs, payload)
		}(payload)
	}
	wg.Wait()
	res := daoWard.InsertMany(ctx, docs)
	if res == -1 {
		daoLog.InsetOne(ctx, model.Log{
			Type:    response.ERROR,
			Local:   "InsertMany-Ward",
			Content: docs,
			From:    "",
			Time:    time.Now().String(),
		})
	}
	return res
}
