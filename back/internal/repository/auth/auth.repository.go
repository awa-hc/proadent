package auth

import (
	"back/internal/domain/entities"
	"context"
)

type AuthRepository interface {
	LoginWithCI(ctx context.Context, email string, password string) (string, error)
	Logout(ctx context.Context, token string) error
	LoginWithEmail(ctx context.Context, email string, password string) (string, error)
	GetUserByContext(ctx context.Context) (*entities.User, error)
}
