package services

import (
	"back/internal/domain/entities"
	"back/internal/repository/user"
	"context"
)

type UserService struct {
	UserRepository user.UserRepository
}

func NewUserService(userRepository user.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (us *UserService) CreateUser(ctx context.Context, user *entities.User) error {
	if err := user.ValidateAtCreate(); err != nil {
		return err
	}

	return us.UserRepository.CreateUser(ctx, user)
}

func (us *UserService) GetByCI(ctx context.Context, ci string) (*entities.User, error) {
	return us.UserRepository.GetByCI(ctx, ci)
}

func (us *UserService) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	return us.UserRepository.GetByEmail(ctx, email)
}

func (us *UserService) GetByUsername(ctx context.Context, username string) (*entities.User, error) {
	return us.UserRepository.GetByUsername(ctx, username)
}

func (us *UserService) UpdateUser(ctx context.Context, user *entities.User) error {
	return us.UserRepository.UpdateUser(ctx, user)
}
func (us *UserService) DeleteUser(ctx context.Context, id int) error {
	return us.UserRepository.DeleteUser(ctx, id)
}
