package menuservice

import (
	"context"
	"time"
	menudomain "vht-go/modules/menu/domain"
	menudtos "vht-go/modules/menu/dtos"

	"github.com/google/uuid"
)

type UpdateMenuCommand struct {
	Id  uuid.UUID
	DTO *menudtos.UpdateMenuDTO
}

type IUpdateMenuRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (*menudomain.Menu, error)
	Update(ctx context.Context, menu *menudomain.Menu, id uuid.UUID) error
}

type UpdateMenuCommandHandler struct {
	repo IUpdateMenuRepository
}

func NewUpdateMenuCommandHandler(repo IUpdateMenuRepository) *UpdateMenuCommandHandler {
	return &UpdateMenuCommandHandler{repo: repo}
}

func (h *UpdateMenuCommandHandler) Handle(ctx context.Context, cmd *UpdateMenuCommand) error {
	if err := cmd.DTO.Validate(); err != nil {
		return err
	}

	// Fetch existing menu
	menu, err := h.repo.FindById(ctx, cmd.Id)
	if err != nil {
		return err
	}

	// Apply updates
	if cmd.DTO.Name != nil {
		menu.Name = *cmd.DTO.Name
	}

	if cmd.DTO.Description != nil {
		menu.Description = cmd.DTO.Description
	}

	if len(cmd.DTO.FoodIds) > 0 {
		menu.FoodIds = menudomain.FoodIDs(cmd.DTO.FoodIds)
	}

	menu.UpdatedAt = time.Now().UTC()

	// Save updates
	if err := h.repo.Update(ctx, menu, cmd.Id); err != nil {
		return err
	}

	return nil
}

