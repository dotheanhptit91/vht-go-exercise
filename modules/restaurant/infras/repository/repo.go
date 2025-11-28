package restaurantrepository

import "gorm.io/gorm"

type GORMRestaurantRepository struct {
	db *gorm.DB
}

func NewGORMRestaurantRepository(db *gorm.DB) *GORMRestaurantRepository {
	return &GORMRestaurantRepository{db: db}
}

