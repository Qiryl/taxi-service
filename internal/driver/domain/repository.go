package domain

import "context"

type Repository interface {
	SignUp(ctx context.Context, d *Driver) error
	SignIn(ctx context.Context, l *Login) (*Driver, error)
}
