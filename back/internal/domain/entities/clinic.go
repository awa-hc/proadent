package entities

import (
	"time"

	"gorm.io/gorm"
)

type Clinic struct {
	gorm.Model
	DateTime        *time.Time `json:"date_time" gorm:"not null"`
	Reason          string     `json:"reason" gorm:"not null"`
	AppointmentCode string     `json:"appointment_code" gorm:"not null;uniqueIndex"`
	DoctorCI        string     `json:"doctor_ci" gorm:"not null"`
	PatientCI       string     `json:"patient_ci" gorm:"not null"`
	VisitDateTime   *time.Time `json:"visit_date_time"`
	Prescription    string     `json:"prescription,omitempty"`
	Notes           string     `json:"notes,omitempty"`
	Code            string     `json:"code" gorm:"uniqueIndex;not null"`
	Status          string     `json:"status" gorm:"not null"`
	Price           float64    `json:"price"`
	Type            string     `json:"type" gorm:"not null"`
}
