package domain

import "context"

type Repository interface {
	SignUp(ctx context.Context, d *Driver) error
	GetDriverByPhone(ctx context.Context, phone string) (*Driver, error)
	GetByPhoneAndEmail(ctx context.Context, phone, email string) (*Driver, error)
}
