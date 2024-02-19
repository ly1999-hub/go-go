package model

import (
	"time"

	"github.com/ly1999-hub/go-go/internal/constant"
	"github.com/ly1999-hub/go-go/internal/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID              primitive.ObjectID `bson:"_id"`
	UserOrder       primitive.ObjectID `bson:"user_order"`
	RestaurantOrder primitive.ObjectID `bson:"restaurant_order"`
	NumberPeople    int16              `bson:"number_people"`
	DishOrder       []DishOrder        `bson:"dish_order"`
	Price           int64              `bson:"price"`
	TimeOrder       string             `bson:"time_order"`
	Action          string             `bson:"action"`
	CreatedAt       string             `bson:"created_at"`
}

type DishOrder struct {
	Dish   Dish
	Number int16
}

type OrderRequest struct {
	UserOrder        string             `json:"user_order"`
	RestaurantOrder  string             `json:"restaurant_order"`
	DishOrderRequest []DishOrderRequest `json:"dish_order_request"`
	NumberPeople     int16              `json:"number_people"`
	TimeOrder        string             `json:"time_order"`
}

type DishOrderRequest struct {
	DishOrderId string `json:"dish_order_id"`
	Number      int16  `json:"number"`
}

func (m OrderRequest) ToOrder() Order {
	return Order{
		ID:              primitive.NewObjectID(),
		UserOrder:       util.ObjectIDFromHex(m.UserOrder),
		RestaurantOrder: util.ObjectIDFromHex(m.RestaurantOrder),
		NumberPeople:    m.NumberPeople,
		Action:          constant.StartOrder,
		TimeOrder:       m.TimeOrder,
		CreatedAt:       time.Now().String(),
	}
}
