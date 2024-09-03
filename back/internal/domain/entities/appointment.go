package entities

import (
	"back/internal/utils"
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	DoctorCI         string             `json:"doctor_ci" gorm:"not null"`
	PatientCI        string             `json:"patient_ci" gorm:"not null"`
	DateTime         *time.Time         `json:"date_time" gorm:"not null"`
	Reason           string             `json:"reason" gorm:"not null"`
	Status           string             `json:"status" gorm:"not null"`
	RequestedBy      string             `json:"requested_by" gorm:"not null"`
	Type             string             `json:"type" gorm:"not null"`
	Code             string             `json:"code" gorm:"uniqueIndex;not null"`
	UserAppointments []UserAppointments `gorm:"foreignKey:AppointmentCode;references:Code"`
}

func (a *Appointment) BeforeCreate(tx *gorm.DB) (err error) {
	a.Status = "pending"
	return
}

func (a *Appointment) BeforeCancelled(tx *gorm.DB) (err error) {
	a.Status = "cancelled"
	return
}

func (a *Appointment) BeforeAccepted(tx *gorm.DB) (err error) {
	a.Status = "accepted"
	return
}

func (a *Appointment) BeforeRejected(tx *gorm.DB) (err error) {
	a.Status = "rejected"
	return
}

func (a *Appointment) BeforeDelete(tx *gorm.DB) (err error) {
	a.Status = "deleted"
	return
}
func (a *Appointment) BeforeFinish(tx *gorm.DB) (err error) {
	a.Status = "finished"
	return
}

func (a *Appointment) ValidateAtCreate() error {

	if err := a.ValidateDateTime(); err != nil {
		return err
	}

	if err := a.ValidateReason(); err != nil {
		return err
	}

	if err := a.ValidateRequestedBy(); err != nil {
		return err
	}

	if err := a.ValidateType(); err != nil {
		return err
	}

	return nil

}

func (a *Appointment) ValidateDateTime() error {
	if a.DateTime.String() == "" {
		return &utils.ValidationError{Field: "date_time", Message: "date_time is required"}
	}
	if a.DateTime.String() < "2024-01-01" {
		return &utils.ValidationError{Field: "date_time", Message: "date_time is invalid"}
	}

	return nil
}
func (a *Appointment) ValidateReason() error {
	if a.Reason == "" {
		return &utils.ValidationError{Field: "reason", Message: "reason is required"}
	}
	if len(a.Reason) < 10 {
		return &utils.ValidationError{Field: "reason", Message: "reason is too short"}
	}
	return nil
}

func (a *Appointment) ValidateRequestedBy() error {
	if a.RequestedBy == "" {
		return &utils.ValidationError{Field: "requested_by", Message: "requested_by is required"}
	}

	return nil
}
func (a *Appointment) ValidateType() error {
	if a.Type == "" {
		return &utils.ValidationError{Field: "type", Message: "type is required"}
	}
	if a.Type != "virtual" && a.Type != "presential" {
		return &utils.ValidationError{Field: "type", Message: "type is invalid"}
	}

	return nil
}

func (a *Appointment) GenerateCode(code string) {
	a.Code = code
}
