package rstlikerepository

import (
	"context"
	rstlikedomain "vht-go/modules/restaurantlike/domain"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type gormRestaurantRepository struct {
	db *gorm.DB
}

func NewGORMRestaurantRepository(db *gorm.DB) *gormRestaurantRepository {
	return &gormRestaurantRepository{db: db}
}

func (r *gormRestaurantRepository) GetRestaurant(ctx context.Context, id int) (*rstlikedomain.Restaurant, error) {
	var restaurant rstlikedomain.Restaurant

	if err := r.db.Where("id = ?", id).First(&restaurant).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return &restaurant, nil
}