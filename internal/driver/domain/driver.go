package domain

import (
	"time"

	"github.com/google/uuid"
)

type Driver struct {
	ID        uuid.UUID
	Name      string
	Phone     string
	Email     string
	Password  string
	TaxiType  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Login struct {
	Phone    string
	Password string
}
