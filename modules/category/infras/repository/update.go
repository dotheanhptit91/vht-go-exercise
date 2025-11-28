package categoryrepository

import (
	"context"
	categorydomain "vht-go/modules/category/domain"

	"github.com/google/uuid"
)

func (repo *GORMCategoryRepository) Update(ctx context.Context, category *categorydomain.Category, id *uuid.UUID) error {
	if err := repo.db.Where("id = ?", *id).Save(category).Error; err != nil {
		return err
	}
	return nil
}