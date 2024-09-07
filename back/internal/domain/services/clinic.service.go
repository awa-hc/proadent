package services

import (
	"back/internal/domain/entities"
	"back/internal/repository/appointment"
	"back/internal/repository/clinic"
	"back/internal/repository/user"
	"context"
)

type ClinicService struct {
	AppointmentRepository appointment.AppointmentRepository
	ClinicRepository      clinic.ClinicRepository
	UserRepository        user.UserRepository
}

func NewClinicService(
	appointmentRepository appointment.AppointmentRepository,
	clinicRepository clinic.ClinicRepository,
	userRepository user.UserRepository,
) *ClinicService {
	return &ClinicService{
		AppointmentRepository: appointmentRepository,
		ClinicRepository:      clinicRepository,
		UserRepository:        userRepository,
	}
}

func (cs *ClinicService) Create(ctx context.Context, clinic *entities.Clinic) error {
	return cs.ClinicRepository.Created(ctx, clinic)
}

func (cs *ClinicService) GetByID(ctx context.Context, id int) (*entities.Clinic, error) {
	return cs.ClinicRepository.GetByID(ctx, id)
}

func (cs *ClinicService) GetByCode(ctx context.Context, code string) (*entities.Clinic, error) {
	return cs.ClinicRepository.GetByCode(ctx, code)
}

func (cs *ClinicService) GetAll(ctx context.Context) ([]entities.Clinic, error) {
	return cs.ClinicRepository.GetAll(ctx)
}

func (cs *ClinicService) GetLast(ctx context.Context) (*entities.Clinic, error) {
	return cs.ClinicRepository.GetLast(ctx)
}

func (cs *ClinicService) Confirm(ctx context.Context, code string) error {
	return cs.ClinicRepository.Confirm(ctx, code)
}

func (cs *ClinicService) Cancel(ctx context.Context, code string) error {
	return cs.ClinicRepository.Cancel(ctx, code)
}

func (cs *ClinicService) Finish(ctx context.Context, code string) error {
	return cs.ClinicRepository.Finish(ctx, code)
}

func (cs *ClinicService) Accept(ctx context.Context, code string) error {
	return cs.ClinicRepository.Accept(ctx, code)
}

func (cs *ClinicService) Reject(ctx context.Context, code string) error {
	return cs.ClinicRepository.Reject(ctx, code)
}

func (cs *ClinicService) Delete(ctx context.Context, code string) error {
	return cs.ClinicRepository.Delete(ctx, code)
}

func (cs *ClinicService) UpdateDoctorCI(ctx context.Context, code string, doctorCI string) error {
	return cs.ClinicRepository.UpdateDoctorCI(ctx, code, doctorCI)
}

func (cs *ClinicService) UpdateDateTime(ctx context.Context, code string, dateTime string) error {
	return cs.ClinicRepository.UpdateDateTime(ctx, code, dateTime)
}

func (cs *ClinicService) GetLastByUserCI(ctx context.Context, userCI string) (*entities.Clinic, error) {
	return cs.ClinicRepository.GetLastByUserCI(ctx, userCI)
}

func (cs *ClinicService) GetLastByDoctorCI(ctx context.Context, doctorCI string) (*entities.Clinic, error) {
	return cs.ClinicRepository.GetLastByDoctorCI(ctx, doctorCI)
}

func (cs *ClinicService) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return cs.ClinicRepository.WithTransaction(ctx, fn)
}

func (cs *ClinicService) GetByDoctorCI(ctx context.Context, doctorCI string) ([]entities.Clinic, error) {
	return cs.ClinicRepository.GetByDoctorCI(ctx, doctorCI)
}

func (cs *ClinicService) GetLastByCI(ctx context.Context, ci string) (*entities.Clinic, error) {
	return cs.ClinicRepository.GetLastByCI(ctx, ci)
}

func (cs *ClinicService) GetByPatientCI(ctx context.Context, patientCI string) ([]entities.Clinic, error) {
	return cs.ClinicRepository.GetByPatientCI(ctx, patientCI)
}

func (cs *ClinicService) UpdateReason(ctx context.Context, code string, reason string) error {
	return cs.ClinicRepository.UpdateReason(ctx, code, reason)
}

func (cs *ClinicService) UpdateStatus(ctx context.Context, code string, status string) error {
	return cs.ClinicRepository.UpdateStatus(ctx, code, status)
}

func (cs *ClinicService) UpdatePrice(ctx context.Context, code string, price float64) error {
	return cs.ClinicRepository.UpdatePrice(ctx, code, price)
}

func (cs *ClinicService) GenerateCode() (string, error) {
	return cs.ClinicRepository.GenerateCode()
}
