package menudtos

import (
	"errors"
	"strings"
	menudomain "vht-go/modules/menu/domain"
)

type CreateMenuDTO struct {
	RestaurantId int    `json:"restaurant_id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description,omitempty"`
	FoodIds      []int  `json:"food_ids" binding:"required,min=1"`
}

func (dto *CreateMenuDTO) Validate() error {
	dto.Name = strings.TrimSpace(dto.Name)
	dto.Description = strings.TrimSpace(dto.Description)

	if dto.Name == "" {
		return errors.New(menudomain.ErrMenuNameRequired)
	}

	if dto.RestaurantId <= 0 {
		return errors.New(menudomain.ErrInvalidRestaurantId)
	}

	if len(dto.FoodIds) == 0 {
		return errors.New(menudomain.ErrFoodIdsRequired)
	}

	// Validate all food IDs are positive
	for _, foodId := range dto.FoodIds {
		if foodId <= 0 {
			return errors.New(menudomain.ErrInvalidFoodIds)
		}
	}

	return nil
}

