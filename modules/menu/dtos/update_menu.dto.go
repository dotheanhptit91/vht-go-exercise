package menudtos

import (
	"errors"
	"strings"
	menudomain "vht-go/modules/menu/domain"
)

type UpdateMenuDTO struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	FoodIds     []int   `json:"food_ids,omitempty"`
}

func (dto *UpdateMenuDTO) Validate() error {
	if dto.Name != nil {
		trimmed := strings.TrimSpace(*dto.Name)
		dto.Name = &trimmed
		
		if *dto.Name == "" {
			return errors.New(menudomain.ErrMenuNameRequired)
		}
	}

	if dto.Description != nil {
		trimmed := strings.TrimSpace(*dto.Description)
		dto.Description = &trimmed
	}

	// Validate food IDs if provided
	if len(dto.FoodIds) > 0 {
		for _, foodId := range dto.FoodIds {
			if foodId <= 0 {
				return errors.New(menudomain.ErrInvalidFoodIds)
			}
		}
	}

	return nil
}

