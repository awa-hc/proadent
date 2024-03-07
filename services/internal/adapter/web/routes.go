package web

import (
	"services/internal/adapter/email"
	appointment "services/internal/adapter/web/createappointment"
	web "services/internal/adapter/web/verifyaccount"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, emailSender *email.SMTPAdapter) {
	emailHandler := web.NewEmailHandler(emailSender)
	AppointmentCreatedHandler := appointment.SendEmailAppointmentCreated(emailSender)

	router.POST("/account-verification", emailHandler)
	router.POST("/appointment-created", AppointmentCreatedHandler)
}
