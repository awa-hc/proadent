package main

import (
	"back/config/initializers"
	"back/internal/delivery/router"
	"log"
	"time"

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

	router.Run(":8080")

}

func SetupRouter(db *gorm.DB) *gin.Engine {
	route := gin.Default()

	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Auth", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.SetupUserRouter(route, db)
	router.SetupAppointmentRouter(route, db)
	router.SetupUserAppointmentsRouter(route, db)
	router.SetupAuthRouter(route, db)
	router.SetupClinicRouter(route, db)

	return route

}
