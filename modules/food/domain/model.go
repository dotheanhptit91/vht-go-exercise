package fooddomain

import (
	"time"

	"github.com/google/uuid"
)

type Food struct {
	Id           int            `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	RestaurantId int            `json:"restaurant_id" gorm:"column:restaurant_id;"`
	CategoryId   *uuid.UUID     `json:"category_id" gorm:"column:category_id;type:varchar(36)"`
	Name         string         `json:"name" gorm:"column:name;"`
	Description  *string        `json:"description,omitempty" gorm:"column:description;"`
	Price        float64        `json:"price" gorm:"column:price;"`
	Status       int            `json:"status" gorm:"column:status;default:1"`
	CreatedAt    time.Time      `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"column:updated_at;"`
	Restaurant   *FoodRestaurant `json:"restaurant" gorm:"-"`
	Category     *FoodCategory   `json:"category" gorm:"-"`
	
}

func (Food) TableName() string {
	return "foods"
}

type FoodRestaurant struct {
	RestaurantId int   `json:"restaurant_id" gorm:"column:restaurant_id;"`
	Name         string `json:"name" gorm:"column:name;"`
}

func (FoodRestaurant) TableName() string {
	return "restaurants"
}

type FoodCategory struct {
	Id   uuid.UUID `json:"id" gorm:"column:id;"`
	Name string    `json:"name" gorm:"column:name;"`
}

func (FoodCategory) TableName() string {
	return "categories"
}