package middleware

import (
	"back/internal/repository/user"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func RequireAuth(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRepository := user.NewUserRepository(*db)

		tokenString, err := c.Cookie("Auth")
		if err != nil {
			fmt.Println(err, "Error in RequireAuth", err.Error())
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET")), nil
		})

		if err != nil {
			c.JSON(401, gin.H{"message": "Unauthorized Testing... 1"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.JSON(401, gin.H{"message": "Unauthorized Testing 1.2..."})
				c.Abort()
				return
			}

			email := claims["email"].(string)
			user, err := userRepository.GetByEmail(c, email)
			if err != nil || user == nil {
				c.JSON(401, gin.H{"message": "Unauthorized Testing 2..."})
				c.Abort()
				return
			}

			// Optionally, set the user in the context if needed later
			c.Set("user", user)

			c.Next()
		} else {
			c.JSON(401, gin.H{"message": "Unauthorized3"})
			c.Abort()
			return
		}
	}
}
