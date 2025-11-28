package categoryservice

import (
	"context"
	"time"
	categorydomain "vht-go/modules/category/domain"
	categorydtos "vht-go/modules/category/dtos"

	"github.com/google/uuid"
)

type CreateCategoryResultCommand struct {
	DTO *categorydtos.CreateCategoryDTO
}

type ICreateCategoryRepository interface {
	Insert(ctx context.Context, category *categorydomain.Category) error
}

type CreateCategoryResultCommandHandler struct {
	repo ICreateCategoryRepository
}

func NewCreateCategoryResultCommandHandler(repo ICreateCategoryRepository) *CreateCategoryResultCommandHandler {
	return &CreateCategoryResultCommandHandler{repo: repo}
}

func (h *CreateCategoryResultCommandHandler) Handle(ctx context.Context, cmd *CreateCategoryResultCommand) (newId *uuid.UUID, err error) {
	if err := cmd.DTO.Validate(); err != nil {
		return nil, err
	}

	newCatId, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	category := categorydomain.Category{
		Id: 		newCatId,
		Name: 	  cmd.DTO.Name,
		Description: cmd.DTO.Description,
		Status:    1,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if err := h.repo.Insert(ctx, &category); err != nil {
		return nil, err
	}
	
	return &newCatId, nil
}