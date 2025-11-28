package rstlikedomain

import (
	"time"

	"github.com/google/uuid"
)

type RestaurantLike struct {
	RestaurantId int       `json:"restaurantId" gorm:"column:restaurant_id;"`
	UserId       uuid.UUID `json:"userId" gorm:"column:user_id;"`
	CreatedAt    time.Time `json:"createdAt" gorm:"column:created_at"`
}

func (RestaurantLike) TableName() string {
	return "restaurant_likes"
}

type Restaurant struct {
	Id        int        `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	OwnerId   int        `json:"ownerId" gorm:"column:owner_id"`
	Name      string     `json:"name" gorm:"column:name"`
	Status    int        `json:"status" gorm:"column:status"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func (r *Restaurant) IsDeleted() bool {
	return r.Status == 0
}