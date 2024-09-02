package main

import (
	"back/config/initializers"
	"back/internal/delivery/handlers"
	"back/internal/delivery/router"
	"back/internal/domain/services"
	"back/internal/repository/appointment"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {

	if err := initializers.LoadEnv(); err != nil {
		log.Fatal(err)
	}

	db, err := initializers.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	router := SetupRouter(db)

	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "pong"})
	// })

	ApointmentRepository := appointment.NewAppointmentRepository(*db)
	appointmentService := services.NewAppointmentService(ApointmentRepository)
	appointmentHandlers := handlers.NewAppointmentHandler(*appointmentService)

	AppointmentGroup := router.Group("/appointment")
	{
		AppointmentGroup.POST("/create", appointmentHandlers.Create)
		AppointmentGroup.GET("/id/:id", appointmentHandlers.GetByID)
		AppointmentGroup.GET("/code/:code", appointmentHandlers.GetByCode)
		AppointmentGroup.GET("/doctor/:doctorCI", appointmentHandlers.GetByDoctorCI)
		AppointmentGroup.GET("/patient/:patientCI", appointmentHandlers.GetByPatientCI)
		AppointmentGroup.PUT("/confirm/:code", appointmentHandlers.Confirm)
		AppointmentGroup.PUT("/cancel/:code", appointmentHandlers.Cancel)
		AppointmentGroup.PUT("/finish/:code", appointmentHandlers.Finish)
		AppointmentGroup.PUT("/accept/:code", appointmentHandlers.Accept)
		AppointmentGroup.PUT("/reject/:code", appointmentHandlers.Reject)
		AppointmentGroup.PUT("/delete/:code", appointmentHandlers.Delete)
		AppointmentGroup.PUT("/update/doctor/:code/:doctorCI", appointmentHandlers.UpdateDoctorCI)
		AppointmentGroup.PUT("/update/patient/:code/:patientCI", appointmentHandlers.UpdatePatientCI)
		AppointmentGroup.PUT("/update/reason/:code/:reason", appointmentHandlers.UpdateReason)
		AppointmentGroup.PUT("/update/datetime/:code/:dateTime", appointmentHandlers.UpdateDateTime)
		AppointmentGroup.PUT("/update/status/:code/:doctorCI", appointmentHandlers.UpdateStatus)
		AppointmentGroup.PUT("/update/price/:code/:price", appointmentHandlers.UpdatePrice)
		AppointmentGroup.PUT("/update/type/:code/:appointmentType", appointmentHandlers.UpdateType)
	}

	router.Run(":8080")

}

func SetupRouter(db *gorm.DB) *gin.Engine {
	route := gin.Default()

	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATH", "GET", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Auth"},
		AllowCredentials: true,
	}))

	router.SetupUserRouter(route, db)

	return route

}
