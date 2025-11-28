package restaurantdtos

import restaurantdomain "vht-go/modules/restaurant/domain"

type UpdateRestaurantDTO struct {
	Name             *string  `json:"name,omitempty"`
	Addr             *string  `json:"addr,omitempty"`
	CityId           *int     `json:"city_id,omitempty"`
	Lat              *float64 `json:"lat,omitempty"`
	Lng              *float64 `json:"lng,omitempty"`
	ShippingFeePerKm *float64 `json:"shipping_fee_per_km,omitempty"`
	Status           *int     `json:"status,omitempty"`
}

func (UpdateRestaurantDTO) TableName() string {
	return restaurantdomain.Restaurant{}.TableName()
}

