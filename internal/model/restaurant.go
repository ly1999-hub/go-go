package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Restaurant struct {
	ID          primitive.ObjectID   `bson:"_id"`
	Owner       primitive.ObjectID   `bson:"owner"`
	Name        string               `bson:"name"`
	Address     string               `bson:"address"`
	Description string               `bson:"description"`
	Active      bool                 `bson:"active"`
	Videos      []string             `bson:"videos"`
	Images      []string             `bson:"images"`
	Dishs       []primitive.ObjectID `bson:"dishs"`
	CreatedAt   string               `bson:"created_at"`
	UpdatedAt   string               `bson:"updated_at"`
}

type RestaurantDetail struct {
	ID primitive.ObjectID
}

type RestaurantCreate struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

func (r RestaurantCreate) ToRestaurant() Restaurant {
	return Restaurant{
		ID:          primitive.NewObjectID(),
		Name:        r.Name,
		Address:     r.Address,
		Description: r.Description,
		Active:      false,
		Dishs:       []primitive.ObjectID{},
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}
}
