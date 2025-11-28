package restaurantdtos

import (
	"errors"
	"strings"
	restaurantdomain "vht-go/modules/restaurant/domain"
)

type CreateRestaurantDTO struct {
	OwnerId          int      `json:"owner_id" binding:"required"`
	CategoryId       *string  `json:"category_id,omitempty"`
	Name             string   `json:"name" binding:"required"`
	Addr             string   `json:"addr" binding:"required"`
	CityId           *int     `json:"city_id,omitempty"`
	Lat              *float64 `json:"lat,omitempty"`
	Lng              *float64 `json:"lng,omitempty"`
	ShippingFeePerKm float64  `json:"shipping_fee_per_km"`
	Status           int      `json:"status"`
}

func (dto *CreateRestaurantDTO) Validate() error {
	dto.Name = strings.TrimSpace(dto.Name)
	dto.Addr = strings.TrimSpace(dto.Addr)

	if dto.Name == "" {
		return errors.New(restaurantdomain.ErrInvalidRestaurantData)
	}

	if dto.Addr == "" {
		return errors.New(restaurantdomain.ErrInvalidRestaurantData)
	}

	// Set default status if not provided
	if dto.Status == 0 {
		dto.Status = 1
	}

	// Ensure shipping fee is not negative
	if dto.ShippingFeePerKm < 0 {
		dto.ShippingFeePerKm = 0
	}

	return nil
}

