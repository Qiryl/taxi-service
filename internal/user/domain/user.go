package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `db:"id"`
	Name         string    `db:"name"`
	Phone        string    `db:"phone"`
	Email        string    `db:"email"`
	Password     string    `db:"password"`
	RegisteredAt time.Time `db:"registration_date"`
}

type LoginRequest struct {
	Phone    string
	Password string
}

type UserUsecase interface {
	Register(ctx context.Context, user *User) error
	Login(ctx context.Context, req *LoginRequest) (*User, error)
}

type UserRepository interface {
	Register(ctx context.Context, user *User) error
	GetUserByPhone(ctx context.Context, phone string) (*User, error)
}
