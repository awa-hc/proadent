package entities

import "gorm.io/gorm"

type Clinic struct {
	gorm.Model
	AppointmentID uint        `json:"appointment_id" gorm:"not null"`
	Appointment   Appointment `json:"appointment"`
}
