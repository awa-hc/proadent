package services

import (
	"back/internal/domain/entities"
	"back/internal/repository/appointment"
	"back/internal/repository/clinic"
	"back/internal/repository/userappointments"
	"back/internal/utils"
	"context"

	"gorm.io/gorm"
)

type AppointmentService struct {
	AppointmentRepository      appointment.AppointmentRepository
	userAppointmentsRepository userappointments.UserAppointmentsRepository
	ClinicRepository           clinic.ClinicRepository
}

func NewAppointmentService(
	appointmentRepository appointment.AppointmentRepository,
	userAppointmentsReporitosty userappointments.UserAppointmentsRepository,
	clinicRepository clinic.ClinicRepository,

) *AppointmentService {
	return &AppointmentService{
		userAppointmentsRepository: userAppointmentsReporitosty,
		AppointmentRepository:      appointmentRepository,
		ClinicRepository:           clinicRepository,
	}
}

func (as *AppointmentService) Created(ctx context.Context, appointment *entities.Appointment) error {
	if err := appointment.ValidateAtCreate(); err != nil {
		return err
	}
	var lastAppointment *entities.Appointment
	lastAppointment, err := as.AppointmentRepository.GetLast(ctx)
	if err != nil {
		return err
	}
	if lastAppointment != nil {
		appointment.Code, _ = utils.GenerateCode(lastAppointment.Code)
	} else {
		appointment.Code = "AP-1"

	}

	appointment.CreatedAt = utils.GetCurrentTime()
	appointment.UpdatedAt = utils.GetCurrentTime()

	return as.AppointmentRepository.WithTransaction(ctx, func(tx context.Context) error {
		if err := as.AppointmentRepository.Created(ctx, appointment); err != nil {
			return err
		}

		userAppointments := &entities.UserAppointments{
			UserCI:          appointment.PatientCI,
			AppointmentCode: appointment.Code,
		}

		if err := as.userAppointmentsRepository.Created(ctx, userAppointments); err != nil {
			return err
		}
		return nil

	})

}

func (as *AppointmentService) GetByID(ctx context.Context, id int) (*entities.Appointment, error) {
	return as.AppointmentRepository.GetByID(ctx, id)
}
func (as *AppointmentService) GetByCode(ctx context.Context, code string) (*entities.Appointment, error) {
	return as.AppointmentRepository.GetByCode(ctx, code)
}
func (as *AppointmentService) GetByDoctorCI(ctx context.Context, doctorCI string) ([]entities.Appointment, error) {
	return as.AppointmentRepository.GetByDoctorCI(ctx, doctorCI)
}
func (as *AppointmentService) GetByPatientCI(ctx context.Context, patientCI string) ([]entities.Appointment, error) {
	return as.AppointmentRepository.GetByPatientCI(ctx, patientCI)
}

func (as *AppointmentService) GetLast(ctx context.Context) (*entities.Appointment, error) {
	return as.AppointmentRepository.GetLast(ctx)
}
func (as *AppointmentService) GetAll(ctx context.Context) ([]entities.Appointment, error) {
	return as.AppointmentRepository.GetAll(ctx)
}

func (as *AppointmentService) Confirm(ctx context.Context, code string) error {

	return as.AppointmentRepository.Confirm(ctx, code)
}
func (as *AppointmentService) Cancel(ctx context.Context, code string) error {
	return as.AppointmentRepository.Cancel(ctx, code)
}
func (as *AppointmentService) Finish(ctx context.Context, code string) error {

	var clinic entities.Clinic
	var appointment *entities.Appointment
	var err error

	appointment, err = as.AppointmentRepository.GetByCode(ctx, code)

	if err != nil {
		return err
	}
	if appointment == nil {
		return &utils.ValidationError{Field: "code", Message: "appointment not found"}
	}

	lastClinic, err := as.ClinicRepository.GetLast(ctx)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if lastClinic == nil {
		clinic.Code = "CL-1"
	} else {

		clinic.Code, _ = utils.GenerateCode(lastClinic.Code)
	}

	clinic.CreatedAt = utils.GetCurrentTime()
	clinic.UpdatedAt = utils.GetCurrentTime()

	return as.AppointmentRepository.WithTransaction(ctx, func(tx context.Context) error {

		if err := as.AppointmentRepository.Finish(ctx, code); err != nil {
			return err
		}
		clinic := entities.Clinic{
			DoctorCI:        appointment.DoctorCI,
			PatientCI:       appointment.PatientCI,
			AppointmentCode: appointment.Code,
			DateTime:        appointment.DateTime,
			Reason:          appointment.Reason,
			Status:          "pending",
			Type:            appointment.Type,
		}

		if err := as.ClinicRepository.Created(ctx, &clinic); err != nil {
			return err
		}

		return nil

	})

}
func (as *AppointmentService) Accept(ctx context.Context, code string) error {
	return as.AppointmentRepository.Accept(ctx, code)
}
func (as *AppointmentService) Reject(ctx context.Context, code string) error {
	return as.AppointmentRepository.Reject(ctx, code)
}
func (as *AppointmentService) Delete(ctx context.Context, code string) error {
	return as.AppointmentRepository.Delete(ctx, code)
}
func (as *AppointmentService) UpdateDoctorCI(ctx context.Context, code string, doctorCI string) error {
	return as.AppointmentRepository.UpdateDoctorCI(ctx, code, doctorCI)
}
func (as *AppointmentService) UpdatePatientCI(ctx context.Context, code string, patientCI string) error {
	return as.AppointmentRepository.UpdatePatientCI(ctx, code, patientCI)
}
func (as *AppointmentService) UpdateDateTime(ctx context.Context, code string, dateTime string) error {
	return as.AppointmentRepository.UpdateDateTime(ctx, code, dateTime)
}
func (as *AppointmentService) UpdateReason(ctx context.Context, code string, reason string) error {
	return as.AppointmentRepository.UpdateReason(ctx, code, reason)
}
func (as *AppointmentService) UpdateStatus(ctx context.Context, code string, status string) error {
	return as.AppointmentRepository.UpdateStatus(ctx, code, status)
}
func (as *AppointmentService) UpdatePrice(ctx context.Context, code string, price float64) error {
	return as.AppointmentRepository.UpdatePrice(ctx, code, price)
}
func (as *AppointmentService) UpdateType(ctx context.Context, code string, appointmentType string) error {
	return as.AppointmentRepository.UpdateType(ctx, code, appointmentType)
}
