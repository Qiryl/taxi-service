package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/Qiryl/taxi-service/internal/user/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepo   domain.UserRepository
	ctxTimeout time.Duration
}

var _ domain.UserUsecase = &UserUsecase{}

func NewUserUsecase(userRepo domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (uc *UserUsecase) Register(ctx context.Context, user *domain.User) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return fmt.Errorf("Usecase Register: GenerateFromPassword: %w", err)
	}

	user.ID, user.Password, user.RegisteredAt = uuid.New(), string(password), time.Now()

	err = uc.userRepo.Register(ctx, user)
	if err != nil {
		return fmt.Errorf("Usecase Register: %w", err)
	}

	return nil
}

func (uc *UserUsecase) Login(ctx context.Context, req *domain.LoginRequest) (*domain.User, error) {
	user, err := uc.userRepo.GetUserByPhone(ctx, req.Phone)
	if err != nil {
		return nil, fmt.Errorf("Usecase Login: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("Usecase Login: CompareHashAndPassword: %w", err)
	}

	return user, nil
}
