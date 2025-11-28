package restaurantrepository

import (
	"context"
	"errors"
	restaurantdomain "vht-go/modules/restaurant/domain"

	"gorm.io/gorm"
)

func (repo *GORMRestaurantRepository) FindById(ctx context.Context, id int) (*restaurantdomain.Restaurant, error) {
	var restaurant restaurantdomain.Restaurant
	err := repo.db.WithContext(ctx).Where("id = ?", id).First(&restaurant).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(restaurantdomain.ErrRestaurantNotFound)
		}
		return nil, err
	}
	return &restaurant, nil
}

func (repo *GORMRestaurantRepository) FindAll(ctx context.Context, offset, limit int) ([]restaurantdomain.Restaurant, error) {
	var restaurants []restaurantdomain.Restaurant
	err := repo.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&restaurants).Error
	if err != nil {
		return nil, err
	}
	return restaurants, nil
}

func (repo *GORMRestaurantRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := repo.db.WithContext(ctx).Model(&restaurantdomain.Restaurant{}).Count(&count).Error
	return count, err
}

func (repo *GORMRestaurantRepository) FindWithFilters(ctx context.Context, ownerId *int, cityId *int, status *int, offset, limit int) ([]restaurantdomain.Restaurant, error) {
	var restaurants []restaurantdomain.Restaurant
	query := repo.db.WithContext(ctx)

	if ownerId != nil {
		query = query.Where("owner_id = ?", *ownerId)
	}
	if cityId != nil {
		query = query.Where("city_id = ?", *cityId)
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	err := query.Offset(offset).Limit(limit).Find(&restaurants).Error
	if err != nil {
		return nil, err
	}
	return restaurants, nil
}

func (repo *GORMRestaurantRepository) CountWithFilters(ctx context.Context, ownerId *int, cityId *int, status *int) (int64, error) {
	var count int64
	query := repo.db.WithContext(ctx).Model(&restaurantdomain.Restaurant{})

	if ownerId != nil {
		query = query.Where("owner_id = ?", *ownerId)
	}
	if cityId != nil {
		query = query.Where("city_id = ?", *cityId)
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	err := query.Count(&count).Error
	return count, err
}

