package categoryservice

import (
	"context"
	categorydomain "vht-go/modules/category/domain"
	categorydtos "vht-go/modules/category/dtos"
)

type ListCategoryQuery struct {
	DTO *categorydtos.ListCategoryDTO
}

type IListCategoryQueryRepository interface {
	FindAll(ctx context.Context, dto *categorydtos.ListCategoryDTO) (categories []categorydomain.Category, err error)
}

type ListCategoryQueryHandler struct {
	repo IListCategoryQueryRepository
}

func NewListCategoryQueryHandler(repo IListCategoryQueryRepository) *ListCategoryQueryHandler {
	return &ListCategoryQueryHandler{repo: repo}
}

func (h *ListCategoryQueryHandler) Handle(ctx context.Context, cmd *ListCategoryQuery) (categories []categorydomain.Category, err error) {
	return h.repo.FindAll(ctx, cmd.DTO)
}
