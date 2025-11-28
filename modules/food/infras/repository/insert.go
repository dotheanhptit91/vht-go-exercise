package foodrepository

import (
	"context"
	fooddomain "vht-go/modules/food/domain"
)

func (repo *GORMFoodRepository) Insert(ctx context.Context, food *fooddomain.Food) error {
	return repo.db.WithContext(ctx).Create(food).Error
}

