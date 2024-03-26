package web

import (
	"services/internal/adapter/email"
	contact "services/internal/adapter/web/contactForm"
	appointment "services/internal/adapter/web/createappointment"
	status "services/internal/adapter/web/updateappointment"
	web "services/internal/adapter/web/verifyaccount"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, emailSender *email.SMTPAdapter) {
	emailHandler := web.NewEmailHandler(emailSender)
	AppointmentCreatedHandler := appointment.SendEmailAppointmentCreated(emailSender)
	AppointmentStatusChange := status.SendEmailAppointmentStatusUpdated(emailSender)
	ContactForm := contact.SendContactForm(emailSender)
	router.POST("/account-verification", emailHandler)
	router.POST("/appointment-created", AppointmentCreatedHandler)
	router.POST("/contact-form", ContactForm)
	router.POST("/appointment-status", AppointmentStatusChange)
}
