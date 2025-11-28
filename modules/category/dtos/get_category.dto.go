package categorydtos

import "github.com/google/uuid"

type GetCategoryDTO struct {
	Id *uuid.UUID `json:"id" binding:"required"`
}