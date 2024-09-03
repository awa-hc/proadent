package auth

import (
	"back/internal/domain/entities"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type gormAuthRepository struct {
	db gorm.DB
}

func NewAuthRepository(db gorm.DB) AuthRepository {
	return &gormAuthRepository{db}
}

// Login implements AuthRepository.
func (r *gormAuthRepository) LoginWithCI(ctx context.Context, email string, password string) (string, error) {
	panic("unimplemented")
}

// Logout implements AuthRepository.
func (r *gormAuthRepository) Logout(ctx context.Context, token string) error {
	panic("unimplemented")
}

// Login implements AuthRepository.
func (r *gormAuthRepository) LoginWithEmail(ctx context.Context, email string, password string) (string, error) {
	var user entities.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"ci":    user.CI,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func (r *gormAuthRepository) GetUserByContext(context.Context) (*entities.User, error) {
	var User entities.User
	return &User, nil
}
