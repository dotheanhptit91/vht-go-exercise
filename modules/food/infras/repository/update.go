package foodrepository

import (
	"context"
	fooddomain "vht-go/modules/food/domain"
)

func (repo *GORMFoodRepository) Update(ctx context.Context, food *fooddomain.Food) error {
	return repo.db.WithContext(ctx).Where("id = ?", food.Id).Save(food).Error
}

