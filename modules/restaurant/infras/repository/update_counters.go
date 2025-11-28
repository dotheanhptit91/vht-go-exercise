package restaurantrepository

import (
	"context"
	restaurantdomain "vht-go/modules/restaurant/domain"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (r *GORMRestaurantRepository) IncreaseLikedCount(ctx context.Context, restaurantId int) error {
	if err := r.db.Table(restaurantdomain.Restaurant{}.TableName()).
		Where("id = ?", restaurantId).
		UpdateColumn("liked_count", gorm.Expr("liked_count + 1")).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *GORMRestaurantRepository) DecreaseLikedCount(ctx context.Context, restaurantId int) error {
	if err := r.db.Table(restaurantdomain.Restaurant{}.TableName()).
		Where("id = ?", restaurantId).
		UpdateColumn("liked_count", gorm.Expr("liked_count - 1")).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}