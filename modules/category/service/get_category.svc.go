package categoryservice

import (
	"context"
	categorydomain "vht-go/modules/category/domain"
	categorydtos "vht-go/modules/category/dtos"

	"github.com/google/uuid"
)

type GetCategoryQuery struct {
	DTO *categorydtos.GetCategoryDTO
}

type IGetCategoryQueryRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (category *categorydomain.Category, err error)
}

type GetCategoryQueryHandler struct {
	repo IGetCategoryQueryRepository
}

func NewGetCategoryQueryHandler(repo IGetCategoryQueryRepository) *GetCategoryQueryHandler {
	return &GetCategoryQueryHandler{repo: repo}
}

func (h *GetCategoryQueryHandler) Handle(ctx context.Context, cmd *GetCategoryQuery) (category *categorydomain.Category, err error) {
	return h.repo.FindById(ctx, *cmd.DTO.Id)
}