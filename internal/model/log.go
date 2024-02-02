package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
	ID      primitive.ObjectID `bson:"_id"`
	Type    string             `bson:"type"`
	Local   string             `bson:"local"`
	Content interface{}        `bson:"content"`
	From    string             `bson:"from"`
	Time    string             `bson:"time"`
}
