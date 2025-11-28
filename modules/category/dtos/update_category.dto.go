package categorydtos

import categorydomain "vht-go/modules/category/domain"

type UpdateCategoryDTO struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Status      *int    `json:"status,omitempty"`
}

func (UpdateCategoryDTO) TableName() string {
	return categorydomain.Category{}.TableName()
}
