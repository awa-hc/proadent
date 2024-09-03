package user

import (
	"back/internal/domain/dto"
	"back/internal/domain/entities"
	"back/internal/utils"
	"context"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type gormUserRepository struct {
	db gorm.DB
}

func NewUserRepository(db gorm.DB) UserRepository {
	return &gormUserRepository{db}
}

// CreateUser implements UserRepository.
func (g *gormUserRepository) CreateUser(ctx context.Context, user *entities.User) error {

	if err := g.db.First(&user, "email = ?", user.Email).Error; err == nil {
		return &utils.ValidationError{Field: "email", Message: "email already exists"}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return g.db.Create(&user).Error
}

func (g *gormUserRepository) GetAll(ctx context.Context) ([]entities.User, error) {
	var users []entities.User
	if err := g.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByCI implements UserRepository. Finds User BY CI number 'string;
func (g *gormUserRepository) GetByCI(ctx context.Context, ci string) (*dto.UserDTO, error) {
	var user entities.User

	if err := g.db.WithContext(ctx).Where("ci = ?", ci).First(&user).Error; err != nil {
		return nil, err
	}

	userDTO := &dto.UserDTO{
		Username:     user.Username,
		Email:        user.Email,
		CI:           user.CI,
		Role:         user.Role,
		Birthdate:    user.Birthdate,
		Appointments: user.Appointments,
	}

	return userDTO, nil

}

// GetUserByEmail implements UserRepository. Finds User by email 'string'.
func (g *gormUserRepository) GetByEmail(ctx context.Context, email string) (*dto.UserDTO, error) {
	var user entities.User

	if err := g.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	userDTO := &dto.UserDTO{
		Username:     user.Username,
		Email:        user.Email,
		CI:           user.CI,
		Role:         user.Role,
		Birthdate:    user.Birthdate,
		Appointments: user.Appointments,
	}
	return userDTO, nil
}

// GetUserByUsername implements UserRepository. Finds User by username 'string'.
func (g *gormUserRepository) GetByUsername(ctx context.Context, username string) (*dto.UserDTO, error) {
	var user entities.User
	if err := g.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	userDTO := &dto.UserDTO{
		Username:     user.Username,
		Email:        user.Email,
		CI:           user.CI,
		Role:         user.Role,
		Birthdate:    user.Birthdate,
		Appointments: user.Appointments,
	}
	return userDTO, nil
}

// UpdateUser implements UserRepository. Updates User.
func (g *gormUserRepository) UpdateUser(ctx context.Context, users *entities.User) error {

	if err := g.db.First(&entities.User{}, users.ID).Error; err != nil {
		return err
	}
	return g.db.Save(&users).Error
}

func (g *gormUserRepository) DeleteUser(ctx context.Context, id int) error {
	if err := g.db.First(&entities.User{}, id).Error; err != nil {
		return err
	}
	return g.db.Delete(&entities.User{}, id).Error
}
