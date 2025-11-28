package userrepository

import "gorm.io/gorm"

type gormUserRepository struct {
	db *gorm.DB
}

func NewGORMUserRepository(db *gorm.DB) *gormUserRepository {
	return &gormUserRepository{db: db}
}