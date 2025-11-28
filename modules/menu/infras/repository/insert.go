package menurepository

import (
	"context"
	menudomain "vht-go/modules/menu/domain"
)

func (repo *GORMMenuRepository) Insert(ctx context.Context, menu *menudomain.Menu) error {
	return repo.db.WithContext(ctx).Create(menu).Error
}

