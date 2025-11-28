package categoryrepository

import (
	"context"
	categorydomain "vht-go/modules/category/domain"

	"github.com/google/uuid"
)

func (repo *GORMCategoryRepository) Delete(ctx context.Context, id *uuid.UUID) error {
	if err := repo.db.Table(categorydomain.Category{}.TableName()).Delete(nil, "id = ?", *id).Error; err != nil {
		return err
	}
	return nil
}