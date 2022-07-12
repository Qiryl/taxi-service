package psql

import (
	"context"
	"fmt"

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
		return nil, err
	}
	return conn, nil
}

func (repo *psqlUserRepo) Register(ctx context.Context, user *domain.User) error {
	//TODO: move connection to struct
	conn, err := repo.psqlConnect(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	query := "INSERT INTO users (name, phone, email, password) VALUES ($1, $2, $3, $4)"

	_, err = repo.db.ExecContext(ctx, query, user.Name, user.Phone, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("Repo Register: failed executing db insert: %w", err)
	}

	return nil
}

func (repo *psqlUserRepo) GetPassByPhone(ctx context.Context, phone string) (string, error) {
	//TODO: move connection to struct
	conn, err := repo.psqlConnect(ctx)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	query := "SELECT password FROM users WHERE phone = $1"
	var password string

	err = repo.db.GetContext(ctx, &password, query, phone)
	if err != nil {
		return "", fmt.Errorf("Repo GetPassByPhone: failed executing db select: %w", err)
	}

	return password, nil
}
