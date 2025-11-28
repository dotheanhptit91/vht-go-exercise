package categorydomain

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	Id          uuid.UUID `json:"id" gorm:"column:id;"`
	Name        string    `json:"name" gorm:"column:name;"`
	Description string    `json:"description" gorm:"column:description;"`
	Status      int       `json:"status" gorm:"column:status;"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;"`
	// Icon        string `json:"icon"`
}

func (Category) TableName() string {
	return "categories"
}
