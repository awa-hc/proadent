package services

import (
	"back/internal/domain/dto"
	"back/internal/domain/entities"
	"back/internal/repository/appointment"
	"back/internal/repository/userappointments"
	"back/internal/utils"
	"context"
)

type UserAppointmentsService struct {
	UserService             *UserService
	AppointmentRepository   appointment.AppointmentRepository
	UserAppointmentsService userappointments.UserAppointmentsRepository
}

func NewUserAppointmentsService(
	userService *UserService,
	appointmentRepository appointment.AppointmentRepository,
	userAppointmentsRepository userappointments.UserAppointmentsRepository,
) *UserAppointmentsService {
	return &UserAppointmentsService{
		UserService:             userService,
		AppointmentRepository:   appointmentRepository,
		UserAppointmentsService: userAppointmentsRepository,
	}
}

func (uas *UserAppointmentsService) Created(ctx context.Context, userAppointments *entities.UserAppointments) error {

	return uas.UserAppointmentsService.Created(ctx, userAppointments)

}

func (uas *UserAppointmentsService) GetByID(ctx context.Context, id int) (*entities.UserAppointments, error) {
	return uas.UserAppointmentsService.GetByID(ctx, id)
}
func (uas *UserAppointmentsService) GetByUserCI(ctx context.Context, userCI string) (dto.UserAppointmentsDTO, error) {

	// validate user exists
	user, err := uas.UserService.GetByCI(ctx, userCI)

	if err != nil {
		return dto.UserAppointmentsDTO{}, err
	}
	if user == nil {
		return dto.UserAppointmentsDTO{}, &utils.ValidationError{Message: "User not found"}
	}

	return uas.UserAppointmentsService.GetByUserCI(ctx, userCI)
}
func (uas *UserAppointmentsService) Delete(ctx context.Context, id int) error {
	return uas.UserAppointmentsService.Delete(ctx, id)
}
