package foodservice

import (
	"context"

	fooddomain "vht-go/modules/food/domain"
)

type DeleteFoodCommand struct {
	Id int
}

type IDeleteFoodRepository interface {
	FindById(ctx context.Context, id int) (*fooddomain.Food, error)
	Delete(ctx context.Context, id int) error
}

type DeleteFoodCommandHandler struct {
	repo IDeleteFoodRepository
}

func NewDeleteFoodCommandHandler(repo IDeleteFoodRepository) *DeleteFoodCommandHandler {
	return &DeleteFoodCommandHandler{repo: repo}
}

func (h *DeleteFoodCommandHandler) Handle(ctx context.Context, cmd *DeleteFoodCommand) error {
	// Check if food exists
	_, err := h.repo.FindById(ctx, cmd.Id)
	if err != nil {
		return err
	}

	return h.repo.Delete(ctx, cmd.Id)
}

