package menurepository

import "gorm.io/gorm"

type GORMMenuRepository struct {
	db *gorm.DB
}

func NewGORMMenuRepository(db *gorm.DB) *GORMMenuRepository {
	return &GORMMenuRepository{db: db}
}

