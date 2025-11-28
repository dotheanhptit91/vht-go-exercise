package categorydtos

import "github.com/google/uuid"

type DeleteCategoryDTO struct {
	Id *uuid.UUID `json:"id" binding:"required"`
}