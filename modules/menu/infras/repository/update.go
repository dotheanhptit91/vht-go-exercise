package menurepository

import (
	"context"
	menudomain "vht-go/modules/menu/domain"

	"github.com/google/uuid"
)

func (repo *GORMMenuRepository) Update(ctx context.Context, menu *menudomain.Menu, id uuid.UUID) error {
	return repo.db.WithContext(ctx).Where("id = ?", id).Save(menu).Error
}

