package restaurantrepository

import (
	"context"
	restaurantdomain "vht-go/modules/restaurant/domain"
)

func (repo *GORMRestaurantRepository) Insert(ctx context.Context, restaurant *restaurantdomain.Restaurant) error {
	return repo.db.WithContext(ctx).Create(restaurant).Error
}

