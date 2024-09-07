package clinic

import (
	"back/internal/domain/entities"
	"back/internal/utils"
	"context"

	"gorm.io/gorm"
)

type gormClinicRepository struct {
	db gorm.DB
}

func NewClinicRepository(db gorm.DB) ClinicRepository {
	return &gormClinicRepository{db}
}

// Accept implements ClinicRepository.
func (g *gormClinicRepository) Accept(ctx context.Context, code string) error {
	return g.db.Model(&entities.Clinic{}).Where("code = ?", code).Update("status", "accepted").Error

}

// Cancel implements ClinicRepository.
func (g *gormClinicRepository) Cancel(ctx context.Context, code string) error {
	return g.db.Model(&entities.Clinic{}).Where("code = ?", code).Update("status", "cancelled").Error
}

// Confirm implements ClinicRepository.
func (g *gormClinicRepository) Confirm(ctx context.Context, code string) error {
	return g.db.Model(&entities.Clinic{}).Where("code = ?", code).Update("status", "confirmed").Error
}

// Created implements ClinicRepository.
func (g *gormClinicRepository) Created(ctx context.Context, clinic *entities.Clinic) error {

	if clinic.Code == "" {
		code, err := g.GenerateCode()
		if err != nil {
			return err
		}
		clinic.Code = code
	} else {
		if err := g.db.Where("code = ?", clinic.Code).First(&entities.Clinic{}).Error; err == nil {
			return &utils.ValidationError{Field: "code", Message: "clinic code already exists"}
		}
	}
	return g.db.Create(clinic).Error

}

// Delete implements ClinicRepository.
func (g *gormClinicRepository) Delete(ctx context.Context, code string) error {
	return g.db.Model(&entities.Clinic{}).Where("code = ?", code).Update("status", "deleted").Error
}

// Finish implements ClinicRepository.
func (g *gormClinicRepository) Finish(ctx context.Context, code string) error {
	return g.db.Model(&entities.Clinic{}).Where("code = ?", code).Update("status", "finished").Error
}

// GenerateCode implements ClinicRepository.
func (g *gormClinicRepository) GenerateCode() (string, error) {
	var lastClinic entities.Clinic

	if err := g.db.Order("id DESC").First(&lastClinic).Error; err != nil {
		if gorm.ErrRecordNotFound == err {

			return "CL-1", nil
		}
	}
	newCode, err := utils.GenerateCode(lastClinic.Code)
	if err != nil {
		return "", err
	}
	return newCode, nil
}

// GetAll implements ClinicRepository.
func (g *gormClinicRepository) GetAll(ctx context.Context) ([]entities.Clinic, error) {
	var clinics []entities.Clinic
	if err := g.db.Find(&clinics).Error; err != nil {
		return nil, err
	}
	return clinics, nil

}

// GetByCode implements ClinicRepository.
func (g *gormClinicRepository) GetByCode(ctx context.Context, code string) (*entities.Clinic, error) {
	var clinic entities.Clinic

	if err := g.db.Where("code = ?", code).First(&clinic).Error; err != nil {
		return nil, err
	}
	return &clinic, nil

}

// GetByDoctorCI implements ClinicRepository.
func (g *gormClinicRepository) GetByDoctorCI(ctx context.Context, doctorCI string) ([]entities.Clinic, error) {

	var clinics []entities.Clinic
	if err := g.db.Where("doctor_ci = ?", doctorCI).Find(&clinics).Error; err != nil {
		return nil, err
	}
	return clinics, nil

}

// GetByID implements ClinicRepository.
func (g *gormClinicRepository) GetByID(ctx context.Context, id int) (*entities.Clinic, error) {
	var clinic entities.Clinic

	if err := g.db.Where("id = ?", id).First(&clinic).Error; err != nil {
		return nil, err
	}
	return &clinic, nil

}

// GetByPatientCI implements ClinicRepository.
func (g *gormClinicRepository) GetByPatientCI(ctx context.Context, patientCI string) ([]entities.Clinic, error) {

	var clinics []entities.Clinic
	if err := g.db.Where("patient_ci = ?", patientCI).Find(&clinics).Error; err != nil {
		return nil, err
	}
	return clinics, nil

}

// GetLast implements ClinicRepository.
func (g *gormClinicRepository) GetLast(ctx context.Context) (*entities.Clinic, error) {
	var clinic entities.Clinic

	if err := g.db.Order("id DESC").First(&clinic).Error; err != nil {
		return nil, err
	}
	return &clinic, nil

}

// GetLastByCI implements ClinicRepository.
func (g *gormClinicRepository) GetLastByCI(ctx context.Context, ci string) (*entities.Clinic, error) {
	var clinic entities.Clinic
	if err := g.db.Where("patient_ci = ?", ci).Order("id DESC").First(&clinic).Error; err != nil {
		return nil, err
	}

	return &clinic, nil

}

// GetLastByDoctorCI implements ClinicRepository.
func (g *gormClinicRepository) GetLastByDoctorCI(ctx context.Context, doctorCI string) (*entities.Clinic, error) {

	var clinic entities.Clinic
	if err := g.db.Where("doctor_ci = ?", doctorCI).Order("id DESC").First(&clinic).Error; err != nil {
		return nil, err
	}
	return &clinic, nil

}

// GetLastByUserCI implements ClinicRepository.
func (g *gormClinicRepository) GetLastByUserCI(ctx context.Context, userCI string) (*entities.Clinic, error) {

	var clinic entities.Clinic
	if err := g.db.Where("patient_ci = ?", userCI).Order("id DESC").First(&clinic).Error; err != nil {
		return nil, err
	}
	return &clinic, nil

}

// Reject implements ClinicRepository.
func (g *gormClinicRepository) Reject(ctx context.Context, code string) error {

	return g.db.Model(&entities.Clinic{}).Where("code = ?", code).Update("status", "rejected").Error

}

// UpdateDateTime implements ClinicRepository.
func (g *gormClinicRepository) UpdateDateTime(ctx context.Context, code string, dateTime string) error {

	return g.db.Model(&entities.Clinic{}).Where("code = ?", code).Update("date_time", dateTime).Error

}

// UpdateDoctorCI implements ClinicRepository.
func (g *gormClinicRepository) UpdateDoctorCI(ctx context.Context, code string, doctorCI string) error {

	return g.db.Model(&entities.Clinic{}).Where("code = ?", code).Update("doctor_ci", doctorCI).Error

}

// UpdatePrice implements ClinicRepository.
func (g *gormClinicRepository) UpdatePrice(ctx context.Context, code string, price float64) error {

	return g.db.Model(&entities.Clinic{}).Where("code = ?", code).Update("price", price).Error

}

// UpdateReason implements ClinicRepository.
func (g *gormClinicRepository) UpdateReason(ctx context.Context, code string, reason string) error {

	return g.db.Model(&entities.Clinic{}).Where("code = ?", code).Update("reason", reason).Error

}

// UpdateStatus implements ClinicRepository.
func (g *gormClinicRepository) UpdateStatus(ctx context.Context, code string, status string) error {

	return g.db.Model(&entities.Clinic{}).Where("code = ?", code).Update("status", status).Error

}

// WithTransaction implements ClinicRepository.
func (g *gormClinicRepository) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return g.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(ctx)
	})
}
