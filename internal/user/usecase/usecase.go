package usecase

import (
	"context"
	"errors"
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
		return err
	}

	err := uc.userRepo.Register(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// Should return whole User strcut
func (uc *UserUsecase) Login(ctx context.Context, req *domain.LoginRequest) error {
	password, err := uc.userRepo.GetPassByPhone(ctx, req.Phone)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password))
	if err != nil {
		return errors.New("Incorrect password: " + err.Error())
	}

	return nil
}
