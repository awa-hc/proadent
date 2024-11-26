package dto

import "time"

type AppointmentDTO struct {
	Code      string     `json:"code"`
	DateTime  *time.Time `json:"date_time"`
	Reason    string     `json:"reason"`
	Status    string     `json:"status"`
	PatientCI string     `json:"patient_ci"`
}
