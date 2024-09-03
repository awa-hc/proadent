package router

import (
	"back/internal/delivery/handlers"
	"back/internal/domain/services"
	"back/internal/repository/appointment"
	"back/internal/repository/userappointments"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAppointmentRouter(router *gin.Engine, db *gorm.DB) {

	ApointmentRepository := appointment.NewAppointmentRepository(*db)
	UserAppointmentsRepository := userappointments.NewUserAppointmentsRepository(*db)

	appointmentService := services.NewAppointmentService(ApointmentRepository, UserAppointmentsRepository)
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

}
