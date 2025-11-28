package categoryservice

import (
	"context"
	"errors"
	"strings"
	"time"
	categorydomain "vht-go/modules/category/domain"
	categorydtos "vht-go/modules/category/dtos"

	"github.com/google/uuid"
)

type UpdateCategoryCommand struct {
	Id *uuid.UUID
	DTO *categorydtos.UpdateCategoryDTO
}

type IUpdateCategoryRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (category *categorydomain.Category, err error)
	Update(ctx context.Context, category *categorydomain.Category, id *uuid.UUID) error
}

type UpdateCategoryCommandHandler struct {
	repo IUpdateCategoryRepository
}

func NewUpdateCategoryCommandHandler(repo IUpdateCategoryRepository) *UpdateCategoryCommandHandler {
	return &UpdateCategoryCommandHandler{repo: repo}
}

func (h *UpdateCategoryCommandHandler) Handle(ctx context.Context, cmd *UpdateCategoryCommand) error {
	getCategoryCommand := &categorydtos.GetCategoryDTO{
		Id: cmd.Id,
	}

	oldCategory, err := h.repo.FindById(ctx, *getCategoryCommand.Id)
	if err != nil {
		return err
	}

	category := &categorydomain.Category{
		Id: *cmd.Id,
		Name: oldCategory.Name,
		Description: oldCategory.Description,
		Status: oldCategory.Status,
		CreatedAt: oldCategory.CreatedAt,
		UpdatedAt: time.Now(),
	}
	
	if cmd.DTO.Name != nil {
		category.Name = strings.TrimSpace(*cmd.DTO.Name)
	}
	if cmd.DTO.Description != nil {
		category.Description = strings.TrimSpace(*cmd.DTO.Description)
	}
	if cmd.DTO.Status != nil {
		if *cmd.DTO.Status < 0 || *cmd.DTO.Status > 1 {
			return  errors.New(categorydomain.ErrInvalidStatusFilter)
		}
		if category.Status == 0 {
			return errors.New(categorydomain.ErrCategoryInactive)
		}
		category.Status = *cmd.DTO.Status
	}

	return h.repo.Update(ctx, category, cmd.Id)
}