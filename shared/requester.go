package shared

import "github.com/google/uuid"

type Requester interface {
	Subject() uuid.UUID
	GetRole() string
}

type requester struct {
	userID uuid.UUID
	role   string
}

func NewRequester(userID string) *requester {
	return &requester{userID: uuid.MustParse(userID)}
}

func (r *requester) Subject() uuid.UUID {
	return r.userID
}

func (r *requester) GetRole() string {
	return r.role
}