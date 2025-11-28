package restaurantrepository

import (
	"context"
	restaurantdomain "vht-go/modules/restaurant/domain"
)

func (repo *GORMRestaurantRepository) Delete(ctx context.Context, id int) error {
	return repo.db.WithContext(ctx).Where("id = ?", id).Delete(&restaurantdomain.Restaurant{}).Error
}

