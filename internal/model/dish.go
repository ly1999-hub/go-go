package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dish struct {
	ID          primitive.ObjectID `bson:"_id"`
	Restaurant  primitive.ObjectID `bson:"retaurant"`
	NameDish    string             `bson:"name_dish" json:"name_dish"`
	Images      []string           `bson:"images"`
	Ingredient  []string           `bson:"ingredient" json:"ingredient"`
	Description string             `bson:"description" json:"description"`
	Price       int64              `bson:"price" json:"price"`
	Active      bool               `bson:"active" json:"active"`
	CreatedAt   string             `bson:"created_at"`
	UpdatedAt   string             `bson:"updated_at"`
}

type DishDto struct {
	NameDish    string   `json:"name_dish"`
	Images      []string `json:"images"`
	Ingredient  []string `json:"ingredient"`
	Description string   `json:"description"`
	Price       int64    `json:"price"`
}

type DishCreate struct {
	Payload []DishDto `json:"payload"`
}

func (d DishCreate) ToDish() []Dish {
	var docs []Dish
	for _, dishCreate := range d.Payload {
		dish := Dish{
			ID:          primitive.NewObjectID(),
			NameDish:    dishCreate.NameDish,
			Images:      dishCreate.Images,
			Ingredient:  dishCreate.Ingredient,
			Description: dishCreate.Description,
			Price:       dishCreate.Price,
			Active:      true,
			CreatedAt:   time.Now().String(),
			UpdatedAt:   time.Now().String(),
		}
		docs = append(docs, dish)
	}

	return docs
}
