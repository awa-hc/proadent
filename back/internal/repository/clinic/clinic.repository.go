package clinic

import (
	"back/internal/domain/entities"
	"context"
)

type ClinicRepository interface {
	Created(ctx context.Context, clinic *entities.Clinic) error
	GetByID(ctx context.Context, id int) (*entities.Clinic, error)
	GetByCode(ctx context.Context, code string) (*entities.Clinic, error)
	GetAll(ctx context.Context) ([]entities.Clinic, error)
	GetLast(ctx context.Context) (*entities.Clinic, error)
	Confirm(ctx context.Context, code string) error
	Cancel(ctx context.Context, code string) error
	Finish(ctx context.Context, code string) error
	Accept(ctx context.Context, code string) error
	Reject(ctx context.Context, code string) error
	Delete(ctx context.Context, code string) error
	UpdateDoctorCI(ctx context.Context, code string, doctorCI string) error
	UpdateDateTime(ctx context.Context, code string, dateTime string) error
	GetLastByUserCI(ctx context.Context, userCI string) (*entities.Clinic, error)
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
	GetLastByDoctorCI(ctx context.Context, doctorCI string) (*entities.Clinic, error)
	GetByDoctorCI(ctx context.Context, doctorCI string) ([]entities.Clinic, error)
	GetLastByCI(ctx context.Context, ci string) (*entities.Clinic, error)
	GetByPatientCI(ctx context.Context, patientCI string) ([]entities.Clinic, error)
	UpdateReason(ctx context.Context, code string, reason string) error
	UpdateStatus(ctx context.Context, code string, status string) error
	UpdatePrice(ctx context.Context, code string, price float64) error
	GenerateCode() (string, error)
}
