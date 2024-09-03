package middleware

import (
	"back/internal/domain/dto"
	"back/internal/repository/user"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func Auth(db *gorm.DB) gin.HandlerFunc {
	userRepository := user.NewUserRepository(*db)

	return func(c *gin.Context) {
		AuthHeader := c.GetHeader("Auth")

		if AuthHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}
		parts := strings.Split(AuthHeader, " ")
		if len(parts) != 2 {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}
		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte("SECRET"), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatus(401)
				return
			}

			email := claims["email"].(string)

			user, err := userRepository.GetByEmail(c, email)
			if err != nil {
				c.JSON(401, gin.H{"error": "Error finding user"})
				return
			}
			userDTO := &dto.UserDTO{
				Username: user.Username,
				Email:    user.Email,
				CI:       user.CI,
				Role:     user.Role,
			}

			c.Set("user", userDTO)
			c.Next()
		} else {
			c.AbortWithStatus(401)
			return
		}

	}

}
