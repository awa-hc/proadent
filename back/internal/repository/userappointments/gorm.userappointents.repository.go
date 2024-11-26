package userappointments

import (
	"back/internal/domain/dto"
	"back/internal/domain/entities"
	"context"

	"gorm.io/gorm"
)

type gormUserAppointmentsRepository struct {
	db gorm.DB
}

// Created implements UserAppointmentsRepository.
func (g *gormUserAppointmentsRepository) Created(ctx context.Context, userAppointments *entities.UserAppointments) error {

	return g.db.Create(&userAppointments).Error
}

// Delete implements UserAppointmentsRepository.
func (g *gormUserAppointmentsRepository) Delete(ctx context.Context, id int) error {

	return g.db.Delete(&entities.UserAppointments{}, id).Error
}

// GetByID implements UserAppointmentsRepository.
func (g *gormUserAppointmentsRepository) GetByID(ctx context.Context, id int) (*entities.UserAppointments, error) {
	panic("unimplemented")
}

// GetByUserCI implements UserAppointmentsRepository.
func (g *gormUserAppointmentsRepository) GetByUserCI(ctx context.Context, userCI string) (dto.UserAppointmentsDTO, error) {
	var userAppointments []entities.UserAppointments
	if err := g.db.Where("user_ci = ?", userCI).Preload("Appointment").Preload("User").Find(&userAppointments).Error; err != nil {
		return dto.UserAppointmentsDTO{}, err
	}

	// Asumiendo que solo hay un user_ci para este caso específico
	var userAppointmentsDTO dto.UserAppointmentsDTO

	// Iterar sobre las citas de usuario
	for _, ua := range userAppointments {
		// Si aún no hemos inicializado el DTO, hacerlo
		if userAppointmentsDTO.UserCI == "" {
			userAppointmentsDTO.UserCI = ua.UserCI
			userAppointmentsDTO.Appointments = []dto.AppointmentDTO{}
		}

		// Asegúrate de que estamos agregando citas para el mismo user_ci
		if ua.UserCI == userAppointmentsDTO.UserCI {
			userAppointmentsDTO.Appointments = append(userAppointmentsDTO.Appointments, dto.AppointmentDTO{
				Code:      ua.Appointment.Code,
				Reason:    ua.Appointment.Reason,
				DateTime:  ua.Appointment.DateTime,
				Status:    ua.Appointment.Status,
				PatientCI: ua.User.CI,
			})
		}
	}

	// Ahora userAppointmentsDTO tiene la estructura agrupada correctamente

	return userAppointmentsDTO, nil

}

func NewUserAppointmentsRepository(db gorm.DB) UserAppointmentsRepository {
	return &gormUserAppointmentsRepository{db}
}
