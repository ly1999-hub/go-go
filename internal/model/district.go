package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type District struct {
	ID       primitive.ObjectID `bson:"_id"`
	Province primitive.ObjectID `bson:"province"`
	Name     string             `bson:"name"`
	Pre      string             `bson:"pre"`
	Code     int                `bson:"code"`
}

type DistrictCreateMany struct {
	Districts  []District `json:"districts"`
	IdProvince string     `json:"id_province"`
}
