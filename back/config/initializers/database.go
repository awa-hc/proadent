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
	db.AutoMigrate(&entities.User{})

	return db, nil
}