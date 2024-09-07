package appointment

import (
	"back/internal/domain/entities"
	"context"
)

type AppointmentRepository interface {
	Created(ctx context.Context, appointment *entities.Appointment) error
	GetByID(ctx context.Context, id int) (*entities.Appointment, error)
	GetByCode(ctx context.Context, code string) (*entities.Appointment, error)
	GetByDoctorCI(ctx context.Context, doctorCI string) ([]entities.Appointment, error)
	GetByPatientCI(ctx context.Context, patientCI string) ([]entities.Appointment, error)
	GetAll(ctx context.Context) ([]entities.Appointment, error)
	GetLast(ctx context.Context) (*entities.Appointment, error)
	Confirm(ctx context.Context, code string) error
	Cancel(ctx context.Context, code string) error
	Finish(ctx context.Context, code string) error
	Accept(ctx context.Context, code string) error
	Reject(ctx context.Context, code string) error
	Delete(ctx context.Context, code string) error
	UpdateDoctorCI(ctx context.Context, code string, doctorCI string) error
	UpdatePatientCI(ctx context.Context, code string, patientCI string) error
	UpdateDateTime(ctx context.Context, code string, dateTime string) error
	UpdateReason(ctx context.Context, code string, reason string) error
	UpdateStatus(ctx context.Context, code string, status string) error
	UpdatePrice(ctx context.Context, code string, price float64) error
	UpdateType(ctx context.Context, code string, appointmentType string) error
	GenerateCode() (string, error)
	GetLastAppointmentByCI(ctx context.Context, ci string) (*entities.Appointment, error)
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
