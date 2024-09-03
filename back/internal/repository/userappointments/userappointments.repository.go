package userappointments

import (
	"back/internal/domain/entities"
	"context"
)

type UserAppointmentsRepository interface {
	Created(ctx context.Context, userAppointments *entities.UserAppointments) error
	GetByID(ctx context.Context, id int) (*entities.UserAppointments, error)
	GetByUserCI(ctx context.Context, userCI string) ([]entities.UserAppointments, error)
	Delete(ctx context.Context, id int) error
}
