package userappointments

import (
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
func (g *gormUserAppointmentsRepository) GetByUserCI(ctx context.Context, userCI string) ([]entities.UserAppointments, error) {
	var userAppointments []entities.UserAppointments

	if err := g.db.Where("user_ci = ?", userCI).Find(&userAppointments).Error; err != nil {
		return nil, err
	}
	return userAppointments, nil

}

func NewUserAppointmentsRepository(db gorm.DB) UserAppointmentsRepository {
	return &gormUserAppointmentsRepository{db}
}
