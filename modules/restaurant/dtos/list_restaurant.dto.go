package restaurantdtos

import "vht-go/shared"

type ListRestaurantDTO struct {
	OwnerId *int          `form:"owner_id"`
	CityId  *int          `form:"city_id"`
	Status  *int          `form:"status"`
	Paging  shared.Paging `form:"paging"`
}

