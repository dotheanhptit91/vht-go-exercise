package foodrepository

import (
	"context"
	"errors"

	fooddomain "vht-go/modules/food/domain"
	"gorm.io/gorm"
)

func (repo *GORMFoodRepository) FindById(ctx context.Context, id int) (*fooddomain.Food, error) {
	var food fooddomain.Food
	if err := repo.db.WithContext(ctx).Where("id = ?", id).First(&food).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(fooddomain.ErrFoodNotFound)
		}
		return nil, err
	}
	return &food, nil
}

func (repo *GORMFoodRepository) FindAll(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]fooddomain.Food, error) {
	var foods []fooddomain.Food
	query := repo.db.WithContext(ctx)

	// Apply filters
	if restaurantId, ok := filter["restaurant_id"]; ok && restaurantId != nil {
		query = query.Where("restaurant_id = ?", restaurantId)
	}

	if categoryId, ok := filter["category_id"]; ok && categoryId != nil {
		query = query.Where("category_id = ?", categoryId)
	}

	if status, ok := filter["status"]; ok && status != nil {
		query = query.Where("status = ?", status)
	}

	// Apply pagination
	query = query.Limit(limit).Offset(offset)

	if err := query.Order("created_at DESC").Find(&foods).Error; err != nil {
		return nil, err
	}

	return foods, nil
}

func (repo *GORMFoodRepository) Count(ctx context.Context, filter map[string]interface{}) (int64, error) {
	var count int64
	query := repo.db.WithContext(ctx).Model(&fooddomain.Food{})

	// Apply filters
	if restaurantId, ok := filter["restaurant_id"]; ok && restaurantId != nil {
		query = query.Where("restaurant_id = ?", restaurantId)
	}

	if categoryId, ok := filter["category_id"]; ok && categoryId != nil {
		query = query.Where("category_id = ?", categoryId)
	}

	if status, ok := filter["status"]; ok && status != nil {
		query = query.Where("status = ?", status)
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

