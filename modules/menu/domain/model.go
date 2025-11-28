package menudomain

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// FoodIDs represents a JSON array of food IDs
type FoodIDs []int

// Scan implements sql.Scanner interface for reading from database
func (f *FoodIDs) Scan(value interface{}) error {
	if value == nil {
		*f = FoodIDs{}
		return nil
	}
	
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	
	return json.Unmarshal(bytes, f)
}

// Value implements driver.Valuer interface for writing to database
func (f FoodIDs) Value() (driver.Value, error) {
	if len(f) == 0 {
		return json.Marshal([]int{})
	}
	return json.Marshal(f)
}

type Menu struct {
	Id           uuid.UUID `json:"id" gorm:"column:id;primaryKey;type:varchar(36)"`
	RestaurantId int       `json:"restaurant_id" gorm:"column:restaurant_id;"`
	Name         string    `json:"name" gorm:"column:name;"`
	Description  *string   `json:"description,omitempty" gorm:"column:description;"`
	FoodIds      FoodIDs   `json:"food_ids" gorm:"column:food_ids;type:json"`
	Status       int       `json:"status" gorm:"column:status;default:1"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at;"`
	
	// Relations populated via RPC (not stored in DB)
	Restaurant *MenuRestaurant `json:"restaurant,omitempty" gorm:"-"`
	Foods      []MenuFood      `json:"foods,omitempty" gorm:"-"`
}

func (Menu) TableName() string {
	return "menus"
}

// MenuRestaurant represents restaurant data fetched via RPC
type MenuRestaurant struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// MenuFood represents food data fetched via RPC
type MenuFood struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Price       float64 `json:"price"`
}

