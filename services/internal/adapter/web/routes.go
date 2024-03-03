package web

import (
	"services/internal/adapter/email"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, emailSender *email.SMTPAdapter) {
	emailHandler := NewEmailHandler(emailSender)
	router.POST("/send-email", emailHandler)
}
