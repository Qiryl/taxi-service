package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uuid.UUID
	Name         string
	Phone        string
	Email        string
	Password     string
	RegisteredAt time.Time
}

type LoginRequest struct {
	Phone    string
	Password string
}

type UserUsecase interface {
	Register(ctx context.Context, user *User) error
	Login(ctx context.Context, req *LoginRequest) error
}

type UserRepository interface {
	Register(ctx context.Context, user *User) error
	GetPassByPhone(ctx context.Context, phone string) (string, error)
}

func (u *User) EncryptPassword() error {
	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return fmt.Errorf("EncryptPassword: %w", err)
	}
	u.Password = string(password)
	return nil
}
