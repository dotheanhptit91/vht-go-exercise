package rstlikerepository

import (
	"context"
	rstlikedomain "vht-go/modules/restaurantlike/domain"
	"vht-go/shared"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type gormRestaurantLikeRepository struct {
	db *gorm.DB
}

func NewGORMRestaurantLikeRepository(db *gorm.DB) *gormRestaurantLikeRepository {
	return &gormRestaurantLikeRepository{db: db}
}

func (r *gormRestaurantLikeRepository) GetRestaurantLike(ctx context.Context, restaurantId int, userId uuid.UUID) (*rstlikedomain.RestaurantLike, error) {
	var restaurantLike rstlikedomain.RestaurantLike

	if err := r.db.Where("restaurant_id = ? AND user_id = ?", restaurantId, userId).First(&restaurantLike).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, shared.ErrDataNotFound
		}

		return nil, errors.WithStack(err)
	}

	return &restaurantLike, nil
}

func (r *gormRestaurantLikeRepository) InsertRestaurantLike(ctx context.Context, restaurantLike *rstlikedomain.RestaurantLike) error {
	if err := r.db.Create(restaurantLike).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *gormRestaurantLikeRepository) DeleteRestaurantLike(ctx context.Context, restaurantId int, userId uuid.UUID) error {
	if err := r.db.Table(rstlikedomain.RestaurantLike{}.TableName()).
		Where("restaurant_id = ? AND user_id = ?", restaurantId, userId).Delete(nil).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}