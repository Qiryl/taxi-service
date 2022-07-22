package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Driver struct {
	ID        uuid.UUID
	Name      string `validate:"required,max=30"`
	Phone     string `validate:"required,e164"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required,min=6,max=25"`
	TaxiType  string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (d *Driver) SetId() {
	d.ID = uuid.New()
}

func (d *Driver) SetCreatedAt() {
	d.CreatedAt = time.Now().UTC()
}

func (d *Driver) SetUpdatedAt() {
	d.UpdatedAt = time.Now().UTC()
}

func (d *Driver) EncryptPassword() (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(d.Password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(password), nil
}

func (d *Driver) CheckPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(d.Password), []byte(password)); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

type Login struct {
	Phone    string `validate:"required"`
	Password string `validate:"required,min=6,max=20"`
}

// func (l *Login) Validate(val *validator.Validate) error {
// 	err := val.Struct(l)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }
