package util

import (
	"github.com/ly1999-hub/go-go/internal/util/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ObjectIDFromHex(idString string) primitive.ObjectID {
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		log.Error("Error Convert ObjectIDFromHex ", log.LogData{
			"idString": idString,
			"error":    err.Error(),
		})
	}
	return id
}
