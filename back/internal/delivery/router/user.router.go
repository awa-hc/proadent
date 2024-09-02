package router

import (
	"back/internal/delivery/handlers"
	"back/internal/domain/services"
	"back/internal/repository/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRouter(router *gin.Engine, db *gorm.DB) {

	UserRepository := user.NewUserRepository(*db)
	userService := services.NewUserService(UserRepository)
	userHandlers := handlers.NewUserHandler(*userService)

	UserGroup := router.Group("/user")
	{
		UserGroup.POST("/create", userHandlers.CreateUser)
		UserGroup.GET("/ci/:ci", userHandlers.GetByCI)
		UserGroup.GET("/email/:email", userHandlers.GetByEmail)
	}
}