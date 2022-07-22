package app

import (
	"context"
	"fmt"

	"github.com/Qiryl/taxi-service/internal/driver/domain"
	"github.com/Qiryl/taxi-service/internal/driver/dtos"
	"github.com/go-playground/validator/v10"
)

// Split to the separate usecases?
type Usecase interface {
	SignUp(ctx context.Context, dto *dtos.DriverDTO) error
	SignIn(ctx context.Context, dto *dtos.LoginDTO) (*dtos.DriverDTO, error)
}

type DriverService struct {
	// order service
	repo domain.Repository
	val  *validator.Validate
}

var _ Usecase = &DriverService{}

func NewDriverService(repo domain.Repository) *DriverService {
	return &DriverService{
		repo: repo,
		val:  validator.New(),
	}
}

func (s *DriverService) SignUp(ctx context.Context, dto *dtos.DriverDTO) error {
	driver := dto.ToModel()

	//
	empty, err := s.repo.GetByPhoneAndEmail(ctx, driver.Phone, driver.Email)
	if empty != nil || err != nil {
		return fmt.Errorf("%w", err)
	}

	driver.SetId()
	driver.SetCreatedAt()
	driver.SetUpdatedAt()

	err = s.repo.SignUp(ctx, driver)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (s *DriverService) SignIn(ctx context.Context, dto *dtos.LoginDTO) (*dtos.DriverDTO, error) {
	login := dto.ToModel()

	driver, err := s.repo.GetDriverByPhone(ctx, login.Phone)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	if err := driver.CheckPassword(login.Password); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return dtos.ToDriverDto(driver), nil
}
