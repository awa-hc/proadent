package services

import (
	"back/internal/domain/entities"
	"back/internal/repository/auth"
	"context"
)

type Auth struct {
	AuthRepository auth.AuthRepository
	UserService    *UserService
}

func NewAuthService(authRepository auth.AuthRepository, userService *UserService) *Auth {
	return &Auth{
		AuthRepository: authRepository,
		UserService:    userService,
	}
}

func (a *Auth) LoginWithCI(ctx context.Context, email string, password string) (string, error) {
	return a.AuthRepository.LoginWithCI(ctx, email, password)
}

func (a *Auth) Logout(ctx context.Context, token string) error {
	return a.AuthRepository.Logout(ctx, token)
}
func (a *Auth) LoginWithEmail(ctx context.Context, email string, password string) (string, error) {
	return a.AuthRepository.LoginWithEmail(ctx, email, password)
}

func (a *Auth) GetUserByContext(ctx context.Context) (*entities.User, error) {
	return a.AuthRepository.GetUserByContext(ctx)
}
