package foodservice

import (
	"context"
	"errors"
	"time"

	fooddomain "vht-go/modules/food/domain"
	fooddtos "vht-go/modules/food/dtos"
)

type UpdateFoodCommand struct {
	Id  int
	DTO *fooddtos.UpdateFoodDTO
}

type IUpdateFoodRepository interface {
	FindById(ctx context.Context, id int) (*fooddomain.Food, error)
	Update(ctx context.Context, food *fooddomain.Food) error
}

type UpdateFoodCommandHandler struct {
	repo IUpdateFoodRepository
}

func NewUpdateFoodCommandHandler(repo IUpdateFoodRepository) *UpdateFoodCommandHandler {
	return &UpdateFoodCommandHandler{repo: repo}
}

func (h *UpdateFoodCommandHandler) Handle(ctx context.Context, cmd *UpdateFoodCommand) error {
	if err := cmd.DTO.Validate(); err != nil {
		return err
	}

	// Find existing food
	food, err := h.repo.FindById(ctx, cmd.Id)
	if err != nil {
		return err
	}

	if food == nil {
		return errors.New(fooddomain.ErrFoodNotFound)
	}

	// Apply updates
	if cmd.DTO.CategoryId != nil {
		food.CategoryId = cmd.DTO.CategoryId
	}

	if cmd.DTO.Name != nil {
		food.Name = *cmd.DTO.Name
	}

	if cmd.DTO.Description != nil {
		food.Description = cmd.DTO.Description
	}

	if cmd.DTO.Price != nil {
		food.Price = *cmd.DTO.Price
	}

	if cmd.DTO.Status != nil {
		food.Status = *cmd.DTO.Status
	}

	food.UpdatedAt = time.Now().UTC()

	return h.repo.Update(ctx, food)
}

