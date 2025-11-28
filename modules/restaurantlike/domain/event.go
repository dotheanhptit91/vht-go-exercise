package rstlikedomain

import "github.com/google/uuid"

type RestaurantEvent struct {
	RestaurantId int       `json:"restaurantId"`
	UserId       uuid.UUID `json:"userId"`
}