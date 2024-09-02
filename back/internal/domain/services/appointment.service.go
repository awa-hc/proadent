package services

import (
	"back/internal/domain/entities"
	"back/internal/repository/appointment"
	"context"
)

type AppointmetService struct {
	AppointmentRepository appointment.AppointmentRepository
}

func NewAppointmentService(appointmentRepository appointment.AppointmentRepository) *AppointmetService {
	return &AppointmetService{
		AppointmentRepository: appointmentRepository,
	}
}

func (as *AppointmetService) Created(ctx context.Context, appointment *entities.Appointment) error {
	return as.AppointmentRepository.Created(ctx, appointment)
}

func (as *AppointmetService) GetByID(ctx context.Context, id int) (*entities.Appointment, error) {
	return as.AppointmentRepository.GetByID(ctx, id)
}
func (as *AppointmetService) GetByCode(ctx context.Context, code string) (*entities.Appointment, error) {
	return as.AppointmentRepository.GetByCode(ctx, code)
}
func (as *AppointmetService) GetByDoctorCI(ctx context.Context, doctorCI string) ([]entities.Appointment, error) {
	return as.AppointmentRepository.GetByDoctorCI(ctx, doctorCI)
}
func (as *AppointmetService) GetByPatientCI(ctx context.Context, patientCI string) ([]entities.Appointment, error) {
	return as.AppointmentRepository.GetByPatientCI(ctx, patientCI)
}
func (as *AppointmetService) Confirm(ctx context.Context, code string) error {
	return as.AppointmentRepository.Confirm(ctx, code)
}
func (as *AppointmetService) Cancel(ctx context.Context, code string) error {
	return as.AppointmentRepository.Cancel(ctx, code)
}
func (as *AppointmetService) Finish(ctx context.Context, code string) error {
	return as.AppointmentRepository.Finish(ctx, code)
}
func (as *AppointmetService) Accept(ctx context.Context, code string) error {
	return as.AppointmentRepository.Accept(ctx, code)
}
func (as *AppointmetService) Reject(ctx context.Context, code string) error {
	return as.AppointmentRepository.Reject(ctx, code)
}
func (as *AppointmetService) Delete(ctx context.Context, code string) error {
	return as.AppointmentRepository.Delete(ctx, code)
}
func (as *AppointmetService) UpdateDoctorCI(ctx context.Context, code string, doctorCI string) error {
	return as.AppointmentRepository.UpdateDoctorCI(ctx, code, doctorCI)
}
func (as *AppointmetService) UpdatePatientCI(ctx context.Context, code string, patientCI string) error {
	return as.AppointmentRepository.UpdatePatientCI(ctx, code, patientCI)
}
func (as *AppointmetService) UpdateDateTime(ctx context.Context, code string, dateTime string) error {
	return as.AppointmentRepository.UpdateDateTime(ctx, code, dateTime)
}
func (as *AppointmetService) UpdateReason(ctx context.Context, code string, reason string) error {
	return as.AppointmentRepository.UpdateReason(ctx, code, reason)
}
func (as *AppointmetService) UpdateStatus(ctx context.Context, code string, status string) error {
	return as.AppointmentRepository.UpdateStatus(ctx, code, status)
}
func (as *AppointmetService) UpdatePrice(ctx context.Context, code string, price float64) error {
	return as.AppointmentRepository.UpdatePrice(ctx, code, price)
}
func (as *AppointmetService) UpdateType(ctx context.Context, code string, appointmentType string) error {
	return as.AppointmentRepository.UpdateType(ctx, code, appointmentType)
}
