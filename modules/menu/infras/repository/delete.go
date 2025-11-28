package menurepository

import (
	"context"
	menudomain "vht-go/modules/menu/domain"

	"github.com/google/uuid"
)

func (repo *GORMMenuRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return repo.db.WithContext(ctx).Where("id = ?", id).Delete(&menudomain.Menu{}).Error
}

func (repo *GORMMenuRepository) SoftDelete(ctx context.Context, id uuid.UUID) error {
	return repo.db.WithContext(ctx).Model(&menudomain.Menu{}).Where("id = ?", id).Update("status", 0).Error
}

