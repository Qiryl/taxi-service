package psql

import (
	"context"
	"errors"

	"github.com/Qiryl/taxi-service/internal/user/domain"
	"github.com/jmoiron/sqlx"
)

type psqlUserRepo struct {
	db *sqlx.DB
}

func NewPsqlUserRepo(db *sqlx.DB) *psqlUserRepo {
	return &psqlUserRepo{db: db}
}

func (repo *psqlUserRepo) psqlConnect(ctx context.Context) (*sqlx.Conn, error) {
	conn, err := repo.db.Connx(ctx)
	if err != nil {
		return nil, errors.New("Failed to connect to psql: " + err.Error())
	}
	return conn, nil
}

func (repo *psqlUserRepo) Register(ctx context.Context, user *domain.User) error {
	conn, err := repo.psqlConnect(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	query := "INSERT INTO users (user_name, user_phone, user_email, user_password) VALUES ($1, $2, $3, $4)"

	// Password already encrypted
	_, err = repo.db.ExecContext(ctx, query, user.Name, user.Phone, user.Email, user.Password)
	if err != nil {
		return errors.New("Failed to register: " + err.Error())
	}

	return nil
}

// Select user by phone number and return password
func (repo *psqlUserRepo) GetPassByPhone(ctx context.Context, phone string) (string, error) {
	conn, err := repo.psqlConnect(ctx)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	query := "SELECT user_password FROM users WHERE user_phone = $1"
	var password string

	err = repo.db.GetContext(ctx, &password, query, phone)
	if err != nil {
		return "", errors.New("Failed to login: " + err.Error())
	}

	return password, nil
}
