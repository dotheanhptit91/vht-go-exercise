package fooddtos

import (
	"vht-go/shared"
)

type ListFoodDTO struct {
	RestaurantId *int    `json:"restaurant_id,omitempty" form:"restaurant_id"`
	CategoryId   *string `json:"category_id,omitempty" form:"category_id"`
	Status       *int    `json:"status,omitempty" form:"status"`
	Paging       shared.Paging
}

