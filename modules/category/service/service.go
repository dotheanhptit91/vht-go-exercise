package categoryservice

import (
	"context"
	categorydomain "vht-go/modules/category/domain"
	categorydtos "vht-go/modules/category/dtos"

	"github.com/google/uuid"
)

type ICategoryService interface {
	// CreateNewCategory(ctx context.Context, dto *CreateCategoryDTO) (newId *uuid.UUID, err error)
	// GetAllCategories(ctx context.Context, dto *ListCategoryDTO) (categories []categorydomain.Category, err error)
	// GetCategoryById(ctx context.Context, id *uuid.UUID) (category *categorydomain.Category, err error)
	// DeleteCategory(ctx context.Context, id *uuid.UUID) error
	// UpdateCategory(ctx context.Context, cmd *UpdateCategoryCommand) error
}

type ICategoryRepository interface {
	Insert(ctx context.Context, category *categorydomain.Category) error
	FindAll(ctx context.Context, dto *categorydtos.ListCategoryDTO) (categories []categorydomain.Category, err error)
	FindById(ctx context.Context, dto *categorydtos.GetCategoryDTO) (category *categorydomain.Category, err error)
	Delete(ctx context.Context, id *uuid.UUID) error
	Update(ctx context.Context, category *categorydomain.Category, id *uuid.UUID) error
}

type CategoryService struct {
	repo ICategoryRepository
}

func NewCategoryService(repo ICategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}