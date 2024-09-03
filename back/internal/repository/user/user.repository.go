package user

import (
	"back/internal/domain/dto"
	"back/internal/domain/entities"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entities.User) error
	GetAll(ctx context.Context) ([]entities.User, error)
	GetByCI(ctx context.Context, ci string) (*dto.UserDTO, error)
	GetByUsername(ctx context.Context, username string) (*dto.UserDTO, error)
	GetByEmail(ctx context.Context, email string) (*dto.UserDTO, error)
	UpdateUser(ctx context.Context, user *entities.User) error
	DeleteUser(ctx context.Context, id int) error
}
