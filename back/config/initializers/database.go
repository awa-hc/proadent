package initializers

import (
	"back/internal/domain/entities"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {

	dbURL := os.Getenv("DATABASE_URL_test")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	if err := db.AutoMigrate(&entities.User{}, &entities.Appointment{}, &entities.UserAppointments{}, &entities.Clinic{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, nil
}
