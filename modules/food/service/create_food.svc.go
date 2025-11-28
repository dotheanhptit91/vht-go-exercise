package foodservice

import (
	"context"
	"time"

	fooddomain "vht-go/modules/food/domain"
	fooddtos "vht-go/modules/food/dtos"
)

type CreateFoodResultCommand struct {
	DTO *fooddtos.CreateFoodDTO
}

type ICreateFoodRepository interface {
	Insert(ctx context.Context, food *fooddomain.Food) error
}

type CreateFoodResultCommandHandler struct {
	repo ICreateFoodRepository
}

func NewCreateFoodResultCommandHandler(repo ICreateFoodRepository) *CreateFoodResultCommandHandler {
	return &CreateFoodResultCommandHandler{repo: repo}
}

func (h *CreateFoodResultCommandHandler) Handle(ctx context.Context, cmd *CreateFoodResultCommand) (*int, error) {
	if err := cmd.DTO.Validate(); err != nil {
		return nil, err
	}

	food := fooddomain.Food{
		RestaurantId: cmd.DTO.RestaurantId,
		CategoryId:   cmd.DTO.CategoryId,
		Name:         cmd.DTO.Name,
		Description:  cmd.DTO.Description,
		Price:        cmd.DTO.Price,
		Status:       1,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}

	if err := h.repo.Insert(ctx, &food); err != nil {
		return nil, err
	}

	return &food.Id, nil
}

