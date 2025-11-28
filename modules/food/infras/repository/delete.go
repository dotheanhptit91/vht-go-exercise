package foodrepository

import (
	"context"
	fooddomain "vht-go/modules/food/domain"
)

func (repo *GORMFoodRepository) Delete(ctx context.Context, id int) error {
	return repo.db.WithContext(ctx).Where("id = ?", id).Delete(&fooddomain.Food{}).Error
}

