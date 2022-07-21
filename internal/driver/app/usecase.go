package app

import (
	"context"
	"fmt"

	"github.com/Qiryl/taxi-service/internal/driver/domain"
	"github.com/Qiryl/taxi-service/internal/driver/dtos"
)

// Split to the separate usecases?
type Usecase interface {
	SignUp(ctx context.Context, dto *dtos.DriverDTO) error
	SignIn(ctx context.Context, dto *dtos.LoginDTO) (*dtos.DriverDTO, error)
}

type DriverService struct {
	repo domain.Repository
}

var _ Usecase = &DriverService{}

func NewDriverUsecase(repo domain.Repository) *DriverService {
	return &DriverService{
		repo: repo,
	}
}

func (d *DriverService) SignUp(ctx context.Context, dto *dtos.DriverDTO) error {
	driver := dto.ToModel()

	//
	empty, err := d.repo.GetByPhoneAndEmail(ctx, driver.Phone, driver.Email)
	if empty != nil || err != nil {
		return fmt.Errorf("%w", err)
	}

	driver.SetId()
	driver.SetCreatedAt()
	driver.SetUpdatedAt()

	err = d.repo.SignUp(ctx, driver)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (d *DriverService) SignIn(ctx context.Context, dto *dtos.LoginDTO) (*dtos.DriverDTO, error) {
	login := dto.ToModel()

	driver, err := d.repo.GetDriverByPhone(ctx, login.Phone)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	if err := driver.CheckPassword(login.Password); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return dtos.ToDriverDto(driver), nil
}
