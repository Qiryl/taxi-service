package driver

import (
	"context"

	"github.com/Qiryl/taxi-service/internal/driver/domain"
	"github.com/Qiryl/taxi-service/internal/driver/dtos"
)

// Split to the separate usecases?
type Usecase interface {
	SignUp(ctx context.Context, dto dtos.DriverDTO) error
	SignIn(ctx context.Context, dto dtos.LoginDTO) (*dtos.DriverDTO, error)
}

type Driver struct {
	repo *domain.Repository
}

var _ Usecase = &Driver{}

func (d *Driver) SignUp(ctx context.Context, dto dtos.DriverDTO) error {
	return nil
}

func (d *Driver) SignIn(ctx context.Context, dto dtos.LoginDTO) (*dtos.DriverDTO, error) {
	return nil, nil
}
