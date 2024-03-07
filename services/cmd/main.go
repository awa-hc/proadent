package main

import (
	"services/internal/adapter/email"
	"services/internal/adapter/web"
	"services/pkg/config"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
}

func main() {
	router := gin.Default()
	emailSender := email.ConfigureSMTP()

	web.RegisterRoutes(router, emailSender)

	router.Run(":8080")

}
