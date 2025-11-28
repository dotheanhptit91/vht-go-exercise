package restaurantrepository

import (
	"context"
	restaurantdomain "vht-go/modules/restaurant/domain"
)

func (repo *GORMRestaurantRepository) Update(ctx context.Context, restaurant *restaurantdomain.Restaurant, id int) error {
	return repo.db.WithContext(ctx).Where("id = ?", id).Save(restaurant).Error
}

