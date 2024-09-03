package router

import (
	"back/internal/delivery/handlers"
	"back/internal/domain/services"
	"back/internal/repository/appointment"
	"back/internal/repository/user"
	"back/internal/repository/userappointments"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserAppointmentsRouter(router *gin.Engine, db *gorm.DB) {

	UserAppointmentsRepository := userappointments.NewUserAppointmentsRepository(*db)

	UserRepository := user.NewUserRepository(*db)
	UserService := services.NewUserService(UserRepository)

	AppointmentRepository := appointment.NewAppointmentRepository(*db)

	userAppointmentsService := services.NewUserAppointmentsService(UserService, AppointmentRepository, UserAppointmentsRepository)

	userAppointmentsHandler := handlers.NewUserAppointmentsHandler(*userAppointmentsService)

	router.POST("/userappointments", userAppointmentsHandler.Created)
	router.GET("/userappointments/:id", userAppointmentsHandler.GetByID)
	router.GET("/userappointments/user/:ci", userAppointmentsHandler.GetByUserCI)
	router.DELETE("/userappointments/:id", userAppointmentsHandler.Delete)
}
