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

var _ domain.UserRepository = &psqlUserRepo{}

func NewPsqlUserRepo(db *sqlx.DB) *psqlUserRepo {
	return &psqlUserRepo{db: db}
}

func (repo *psqlUserRepo) Register(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (id, name, phone, email, password, registration_date) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := repo.db.ExecContext(ctx, query, user.ID, user.Name, user.Phone, user.Email, user.Password, user.RegisteredAt)
	if err != nil {
		return fmt.Errorf("Repo Register: failed executing db insert: %w", err)
	}

	return nil
}

func (repo *psqlUserRepo) GetUserByPhone(ctx context.Context, phone string) (*domain.User, error) {
	query := `SELECT * FROM users WHERE phone = $1`
	var user domain.User

	err := repo.db.GetContext(ctx, &user, query, phone)
	if err != nil {
		return nil, fmt.Errorf("Repo GetPassByPhone: failed executing db select: %w", err)
	}

	return &user, nil
}
