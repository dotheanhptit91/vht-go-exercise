package categoryrepository

import (
	"context"
	categorydomain "vht-go/modules/category/domain"
)

func (repo *GORMCategoryRepository) Insert(ctx context.Context, category *categorydomain.Category) error {
	if err := repo.db.Create(category).Error; err != nil {
		return err
	}
	return nil
}