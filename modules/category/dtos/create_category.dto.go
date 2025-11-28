package categorydtos

import (
	"errors"
	"strings"
	categorydomain "vht-go/modules/category/domain"
)

type CreateCategoryDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description,omitempty"`
}

func (dto *CreateCategoryDTO) Validate() error {
	dto.Name = strings.TrimSpace(dto.Name)
	dto.Description = strings.TrimSpace(dto.Description)

	if dto.Name == "" {
		return errors.New(categorydomain.ErrCategoryNameRequired)
	}
	return nil
}