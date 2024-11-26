package userappointments

import (
	"back/internal/domain/dto"
	"back/internal/domain/entities"
	"context"
)

type UserAppointmentsRepository interface {
	Created(ctx context.Context, userAppointments *entities.UserAppointments) error
	GetByID(ctx context.Context, id int) (*entities.UserAppointments, error)
	GetByUserCI(ctx context.Context, userCI string) (dto.UserAppointmentsDTO, error)
	Delete(ctx context.Context, id int) error
}
