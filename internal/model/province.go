package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Province struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
	Pre  string             `bson:"pre"`
	Code int                `bson:"code"`
}

type ProvinceCreateMany struct {
	Provinces []Province
}
