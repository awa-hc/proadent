package router

import (
	"back/internal/delivery/handlers"
	"back/internal/domain/services"
	"back/internal/middleware"
	"back/internal/repository/auth"
	"back/internal/repository/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthRouter(router *gin.Engine, db *gorm.DB) {

	UserRepository := user.NewUserRepository(*db)
	UserService := services.NewUserService(UserRepository)

	AuthRepository := auth.NewAuthRepository(*db)
	authService := services.NewAuthService(AuthRepository, UserService)
	AuthHandler := handlers.NewAuthHandler(*authService)

	AuthGroup := router.Group("/auth")
	{

		AuthGroup.POST("/login", AuthHandler.LoginWithCI)
		AuthGroup.POST("/login/email", AuthHandler.LoginWithEmail)
		AuthGroup.POST("/logout", AuthHandler.Logout)
		AuthGroup.GET("/user", middleware.RequireAuth(db), AuthHandler.GetUserByContext)

	}

	// router.POST("/login", AuthHandler.LoginWithCI)
	// router.POST("/login/email", AuthHandler.LoginWithEmail)
	// router.POST("/logout", AuthHandler.Logout)
	// router.GET("/user", AuthHandler.GetUserByContext)

}
