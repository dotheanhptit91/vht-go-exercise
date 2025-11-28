package menudtos

import "vht-go/shared"

type ListMenuDTO struct {
	RestaurantId *int          `form:"restaurant_id"`
	Status       *int          `form:"status"`
	Paging       shared.Paging `form:"paging"`
}

