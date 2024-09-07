package router

import (
	"back/internal/delivery/handlers"
	"back/internal/domain/services"
	"back/internal/repository/appointment"
	"back/internal/repository/clinic"
	"back/internal/repository/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupClinicRouter(router *gin.Engine, db *gorm.DB) {

	AppointmentRepository := appointment.NewAppointmentRepository(*db)
	UserRepository := user.NewUserRepository(*db)

	ClinicRepository := clinic.NewClinicRepository(*db)
	ClinicService := services.NewClinicService(AppointmentRepository, ClinicRepository, UserRepository)

	ClinicHandlers := handlers.NewClinicHandler(*ClinicService)

	ClinicGroup := router.Group("/clinic")
	{
		ClinicGroup.POST("/create", ClinicHandlers.Create)
		ClinicGroup.GET("/id/:id", ClinicHandlers.GetByID)
		ClinicGroup.GET("/code/:code", ClinicHandlers.GetByCode)
		ClinicGroup.GET("/all", ClinicHandlers.GetAll)
		ClinicGroup.GET("/last", ClinicHandlers.GetLast)
		ClinicGroup.PUT("/confirm/:code", ClinicHandlers.Confirm)
		ClinicGroup.PUT("/cancel/:code", ClinicHandlers.Cancel)
		ClinicGroup.PUT("/finish/:code", ClinicHandlers.Finish)
		ClinicGroup.PUT("/accept/:code", ClinicHandlers.Accept)
		ClinicGroup.PUT("/reject/:code", ClinicHandlers.Reject)
		ClinicGroup.PUT("/delete/:code", ClinicHandlers.Delete)
		ClinicGroup.PUT("/update/doctorCI/:code/:ci", ClinicHandlers.UpdateDoctorCI)
		ClinicGroup.PUT("/update/datetime/:code/:dateTime", ClinicHandlers.UpdateDateTime)
		ClinicGroup.GET("/last/user/:ci", ClinicHandlers.GetLastByUserCI)
		ClinicGroup.GET("/last/doctor/:ci", ClinicHandlers.GetLastByDoctorCI)
		ClinicGroup.GET("/doctor/:ci", ClinicHandlers.GetByDoctorCI)
		ClinicGroup.GET("/last/ci/:ci", ClinicHandlers.GetLastByCI)
		ClinicGroup.GET("/patient/:ci", ClinicHandlers.GetByPatientCI)
		ClinicGroup.PUT("/update/reason/:code/:reason", ClinicHandlers.UpdateReason)
		ClinicGroup.PUT("/update/status/:code/:status", ClinicHandlers.UpdateStatus)
		ClinicGroup.PUT("/update/price/:code/:price", ClinicHandlers.UpdatePrice)

		// ClinicGroup.PUT("/update/reason/:code/:reason", ClinicHandlers.UpdateReason)

	}

}
