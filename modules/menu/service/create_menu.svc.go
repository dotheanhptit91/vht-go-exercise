package menuservice

import (
	"context"
	"time"
	menudomain "vht-go/modules/menu/domain"
	menudtos "vht-go/modules/menu/dtos"

	"github.com/google/uuid"
)

type CreateMenuResultCommand struct {
	DTO *menudtos.CreateMenuDTO
}

type ICreateMenuRepository interface {
	Insert(ctx context.Context, menu *menudomain.Menu) error
}

type CreateMenuResultCommandHandler struct {
	repo ICreateMenuRepository
}

func NewCreateMenuResultCommandHandler(repo ICreateMenuRepository) *CreateMenuResultCommandHandler {
	return &CreateMenuResultCommandHandler{repo: repo}
}

func (h *CreateMenuResultCommandHandler) Handle(ctx context.Context, cmd *CreateMenuResultCommand) (*uuid.UUID, error) {
	if err := cmd.DTO.Validate(); err != nil {
		return nil, err
	}

	newId, _ := uuid.NewV7()
	
	var description *string
	if cmd.DTO.Description != "" {
		description = &cmd.DTO.Description
	}
	
	menu := menudomain.Menu{
		Id:           newId,
		RestaurantId: cmd.DTO.RestaurantId,
		Name:         cmd.DTO.Name,
		Description:  description,
		FoodIds:      menudomain.FoodIDs(cmd.DTO.FoodIds),
		Status:       1,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}

	if err := h.repo.Insert(ctx, &menu); err != nil {
		return nil, err
	}

	return &newId, nil
}

