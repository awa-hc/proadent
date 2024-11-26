package dto

type UserAppointmentsDTO struct {
	UserCI       string           `json:"user_ci"`
	Appointments []AppointmentDTO `json:"appointments"`
}
