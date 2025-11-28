package restaurantdomain

import (
	"time"

	"github.com/google/uuid"
)

type Restaurant struct {
	Id               int                 `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	CategoryId       *uuid.UUID          `json:"categoryId" gorm:"column:category_id"`
	OwnerId          int                 `json:"ownerId" gorm:"column:owner_id"`
	Name             string              `json:"name" gorm:"column:name"`
	Addr             string              `json:"addr" gorm:"column:addr"`
	CityId           *int                `json:"cityId" gorm:"column:city_id"`
	Lat              *float64            `json:"lat" gorm:"column:lat"`
	Lng              *float64            `json:"lng" gorm:"column:lng"`
	Cover            *string             `json:"cover" gorm:"column:cover;type:json"`
	Logo             *string             `json:"logo" gorm:"column:logo;type:json"`
	ShippingFeePerKm float64             `json:"shippingFeePerKm" gorm:"column:shipping_fee_per_km"`
	Status           int                 `json:"status" gorm:"column:status"`
	LikedCount       int                 `json:"likedCount" gorm:"column:liked_count"`
	CreatedAt        *time.Time          `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt        *time.Time          `json:"updatedAt" gorm:"column:updated_at"`
	Category         *RestaurantCategory `json:"category" gorm:"-"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func (r *Restaurant) IsDeleted() bool {
	return r.Status == 0
}

type RestaurantCategory struct {
	Id   uuid.UUID `json:"id" gorm:"column:id;"`
	Name string    `json:"name" gorm:"column:name;"`
}

func (RestaurantCategory) TableName() string {
	return "categories"
}