package shared

import (
	"time"

	"github.com/google/uuid"
)

type PublicUser struct {
	Id        uuid.UUID `json:"id" gorm:"column:id;primaryKey"`
	LastName  string    `json:"lastName" gorm:"column:last_name"`
	FirstName string    `json:"firstName" gorm:"column:first_name"`
	Role      string    `json:"role" gorm:"column:role"`
	Status    int       `json:"status" gorm:"column:status"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (PublicUser) TableName() string {
	return "users"
}