package fooddtos

import (
	"errors"
	"strings"

	fooddomain "vht-go/modules/food/domain"

	"github.com/google/uuid"
)

type CreateFoodDTO struct {
	RestaurantId int        `json:"restaurant_id" binding:"required"`
	CategoryId   *uuid.UUID `json:"category_id,omitempty"`
	Name         string     `json:"name" binding:"required"`
	Description  *string    `json:"description,omitempty"`
	Price        float64    `json:"price" binding:"required"`
}

func (dto *CreateFoodDTO) Validate() error {
	dto.Name = strings.TrimSpace(dto.Name)
	if dto.Description != nil {
		trimmed := strings.TrimSpace(*dto.Description)
		dto.Description = &trimmed
	}

	if dto.Name == "" {
		return errors.New(fooddomain.ErrFoodNameRequired)
	}

	if dto.RestaurantId <= 0 {
		return errors.New(fooddomain.ErrRestaurantIdRequired)
	}

	if dto.Price <= 0 {
		return errors.New(fooddomain.ErrFoodPriceInvalid)
	}

	return nil
}

