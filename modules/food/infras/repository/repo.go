package foodrepository

import "gorm.io/gorm"

type GORMFoodRepository struct {
	db *gorm.DB
}

func NewGORMFoodRepository(db *gorm.DB) *GORMFoodRepository {
	return &GORMFoodRepository{db: db}
}

