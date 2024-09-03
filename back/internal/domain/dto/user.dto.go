package dto

import "back/internal/domain/entities"

type UserDTO struct {
	Username     string                 `json:"username" gorm:"unique;not null"`
	Birthdate    string                 `json:"birthdate"`
	Email        string                 `json:"email" gorm:"unique;not null"`
	Role         string                 `json:"role" gorm:"not null"`
	CI           string                 `json:"ci" gorm:"uniqueIndex;not null"`
	Appointments []entities.Appointment `gorm:"foreignKey:PatientCI;references:CI"`
}
