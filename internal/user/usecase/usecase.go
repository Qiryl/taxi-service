package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/Qiryl/taxi-service/internal/user/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepo   domain.UserRepository
	ctxTimeout time.Duration
}

func NewUserUsecase(userRepo domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo:   userRepo,
		ctxTimeout: 0,
	}
}

func (uc *UserUsecase) Register(ctx context.Context, user *domain.User) error {
	if err := user.EncryptPassword(); err != nil {
		return fmt.Errorf("Usecase Register: %w", err)
	}

	err := uc.userRepo.Register(ctx, user)
	if err != nil {
		return fmt.Errorf("Usecase Register: %w", err)
	}

	return nil
}

// TODO: return whole User strcut
func (uc *UserUsecase) Login(ctx context.Context, req *domain.LoginRequest) error {
	password, err := uc.userRepo.GetPassByPhone(ctx, req.Phone)
	if err != nil {
		return fmt.Errorf("Usecase Login: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password))
	if err != nil {
		return fmt.Errorf("Usecase Login: CompareHashAndPassword: %w", err)
	}

	return nil
}
