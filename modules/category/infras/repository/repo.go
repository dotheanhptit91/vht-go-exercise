package categoryrepository

import "gorm.io/gorm"

type GORMCategoryRepository struct {
	db *gorm.DB
}

func NewGORMCategoryRepository(db *gorm.DB) *GORMCategoryRepository {
	return &GORMCategoryRepository{db: db}
}