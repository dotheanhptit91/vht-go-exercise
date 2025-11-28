package userdomain

import (
	"time"

	"github.com/google/uuid"
)

type UserStatus int

const (
	UserStatusDeleted UserStatus = 0
	UserStatusActive  UserStatus = 1
)

type User struct {
	Id        uuid.UUID    `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Salt      string `json:"salt"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	Status    int    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) IsDeleted() bool {
	return u.Status == int(UserStatusDeleted)
}

func (u *User) Subject() uuid.UUID {
	return u.Id
}

func (u *User) GetRole() string {
	return u.Role
}

func (u *User) Mask() {
	u.Password = ""
	u.Salt = ""
}