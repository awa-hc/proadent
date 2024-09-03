package entities

import "back/internal/utils"

type UserAppointments struct {
	ID              uint        `json:"id" gorm:"primaryKey"`
	UserCI          string      `json:"user_ci"`
	User            User        `gorm:"foreignKey:UserCI;references:CI"`
	AppointmentCode string      `json:"appointment_code"`
	Appointment     Appointment `gorm:"foreignKey:AppointmentCode;references:Code"`
}

func (ua *UserAppointments) ValidateAtCreate() error {
	if err := ua.ValidateUserCI(); err != nil {
		return err
	}

	if err := ua.ValidateAppointmentID(); err != nil {
		return err
	}

	return nil
}

func (ua *UserAppointments) ValidateUserCI() error {
	if ua.UserCI == "" {
		return &utils.ValidationError{Field: "user_ci", Message: "User CI is required"}
	}

	return nil
}

func (ua *UserAppointments) ValidateAppointmentID() error {
	if ua.AppointmentCode == "" {
		return &utils.ValidationError{Field: "appointment_id", Message: "Appointment ID is required"}
	}

	return nil
}
