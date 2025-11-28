package fooddtos

import (
	"errors"
	"strings"

	fooddomain "vht-go/modules/food/domain"

	"github.com/google/uuid"
)

type UpdateFoodDTO struct {
	CategoryId  *uuid.UUID `json:"category_id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	Price       *float64   `json:"price,omitempty"`
	Status      *int       `json:"status,omitempty"`
}

func (dto *UpdateFoodDTO) Validate() error {
	if dto.Name != nil {
		trimmed := strings.TrimSpace(*dto.Name)
		dto.Name = &trimmed
		if *dto.Name == "" {
			return errors.New(fooddomain.ErrFoodNameRequired)
		}
	}

	if dto.Description != nil {
		trimmed := strings.TrimSpace(*dto.Description)
		dto.Description = &trimmed
	}

	if dto.Price != nil && *dto.Price <= 0 {
		return errors.New(fooddomain.ErrFoodPriceInvalid)
	}

	return nil
}

