package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ward struct {
	ID       primitive.ObjectID `bson:"_id"`
	District primitive.ObjectID `bson:"district"`
	Name     string             `bson:"name"`
	Pre      string             `bson:"pre"`
}

type WardCreateMany struct {
	Wards      []Ward
	IdDistrict string
}
