package appointment

import (
	"back/internal/domain/entities"
	"back/internal/utils"
	"context"

	"gorm.io/gorm"
)

type gormAppointmentRepository struct {
	db gorm.DB
}

func NewAppointmentRepository(db gorm.DB) AppointmentRepository {
	return &gormAppointmentRepository{db}
}

// Accept implements AppointmentRepository.
func (g *gormAppointmentRepository) Accept(ctx context.Context, code string) error {
	if err := g.db.First(&entities.Appointment{}, code).Error; err != nil {
		return err
	}
	return g.db.Model(&entities.Appointment{}).Where("code = ?", code).Update("status", "accepted").Error

}

// Cancel implements AppointmentRepository.
func (g *gormAppointmentRepository) Cancel(ctx context.Context, code string) error {
	if err := g.db.First(&entities.Appointment{}, code).Error; err != nil {
		return err
	}
	return g.db.Model(&entities.Appointment{}).Where("code = ?", code).Update("status", "cancelled").Error

}

// Confirm implements AppointmentRepository.
func (g *gormAppointmentRepository) Confirm(ctx context.Context, code string) error {

	return g.db.Model(&entities.Appointment{}).Where("code = ?", code).Update("status", "confirmed").Error

}

// Createt implements AppointmentRepository.
func (g *gormAppointmentRepository) Created(ctx context.Context, appointment *entities.Appointment) error {

	if appointment.Code == "" {
		code, err := g.GenerateCode()
		if err != nil {
			return err
		}
		appointment.Code = code
	} else {
		if err := g.db.Where("code = ?", appointment.Code).First(&entities.Appointment{}).Error; err == nil {
			return &utils.ValidationError{Field: "code", Message: "appointment code already exists"}
		}
	}

	if err := g.db.Create(appointment).Error; err != nil {
		return err
	}
	return nil

}

// Delete implements AppointmentRepository.
func (g *gormAppointmentRepository) Delete(ctx context.Context, code string) error {
	return g.db.Model(&entities.Appointment{}).Where("code = ?", code).Update("status", "deleted").Error
}

// Finish implements AppointmentRepository.
func (g *gormAppointmentRepository) Finish(ctx context.Context, code string) error {

	return g.db.Model(&entities.Appointment{}).Where("code = ?", code).Update("status", "finished").Error

}

// GetByCode implements AppointmentRepository.
func (g *gormAppointmentRepository) GetByCode(ctx context.Context, code string) (*entities.Appointment, error) {

	var appointment entities.Appointment
	if err := g.db.Where("code = ?", code).First(&appointment).Preload("ClinicVisit").Error; err != nil {
		return nil, err
	}

	return &appointment, nil

}

func (g *gormAppointmentRepository) GetLastAppointmentByCI(ctx context.Context, ci string) (*entities.Appointment, error) {
	var appointment entities.Appointment
	if err := g.db.Where("patient_ci = ?", ci).Order("id DESC").First(&appointment).Error; err != nil {
		return nil, err
	}
	return &appointment, nil
}

// GetByDoctorCI implements AppointmentRepository.
func (g *gormAppointmentRepository) GetByDoctorCI(ctx context.Context, doctorCI string) ([]entities.Appointment, error) {

	var appointments []entities.Appointment
	if err := g.db.Where("doctor_ci = ?", doctorCI).Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil

}

// GetByID implements AppointmentRepository.
func (g *gormAppointmentRepository) GetByID(ctx context.Context, id int) (*entities.Appointment, error) {
	if err := g.db.First(&entities.Appointment{}, id).Error; err != nil {
		return nil, err
	}
	var appointment entities.Appointment
	if err := g.db.Where("id = ?", id).First(&appointment).Error; err != nil {
		return nil, err
	}
	return &appointment, nil
}

// GetByPatientCI implements AppointmentRepository.
func (g *gormAppointmentRepository) GetByPatientCI(ctx context.Context, patientCI string) ([]entities.Appointment, error) {

	var appointments []entities.Appointment
	if err := g.db.Where("patient_ci = ?", patientCI).Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil

}

func (g *gormAppointmentRepository) GetLast(ctx context.Context) (*entities.Appointment, error) {
	var appointment entities.Appointment
	if err := g.db.Order("id DESC").First(&appointment).Error; err != nil {
		return nil, err
	}
	return &appointment, nil
}
func (g *gormAppointmentRepository) GetAll(ctx context.Context) ([]entities.Appointment, error) {
	var appointments []entities.Appointment
	if err := g.db.Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}

// Reject implements AppointmentRepository.
func (g *gormAppointmentRepository) Reject(ctx context.Context, code string) error {
	if err := g.db.First(&entities.Appointment{}, code).Error; err != nil {
		return err
	}
	return g.db.Model(&entities.Appointment{}).Where("code = ?", code).Update("status", "rejected").Error

}

// UpdateDateTime implements AppointmentRepository.
func (g *gormAppointmentRepository) UpdateDateTime(ctx context.Context, code string, dateTime string) error {
	if err := g.db.First(&entities.Appointment{}, code).Error; err != nil {
		return err
	}
	return g.db.Model(&entities.Appointment{}).Where("code = ?", code).Update("date_time", dateTime).Error

}

// UpdateDoctorCI implements AppointmentRepository.
func (g *gormAppointmentRepository) UpdateDoctorCI(ctx context.Context, code string, doctorCI string) error {
	if err := g.db.First(&entities.Appointment{}, code).Error; err != nil {
		return err
	}
	return g.db.Model(&entities.Appointment{}).Where("code = ?", code).Update("doctor_ci", doctorCI).Error

}

// UpdatePatientCI implements AppointmentRepository.
func (g *gormAppointmentRepository) UpdatePatientCI(ctx context.Context, code string, patientCI string) error {
	if err := g.db.First(&entities.Appointment{}, code).Error; err != nil {
		return err
	}
	return g.db.Model(&entities.Appointment{}).Where("code = ?", code).Update("patient_ci", patientCI).Error

}

// UpdatePrice implements AppointmentRepository.
func (g *gormAppointmentRepository) UpdatePrice(ctx context.Context, code string, price float64) error {
	if err := g.db.First(&entities.Appointment{}, code).Error; err != nil {
		return err
	}
	return g.db.Model(&entities.Appointment{}).Where("code = ?", code).Update("price", price).Error

}

// UpdateReason implements AppointmentRepository.
func (g *gormAppointmentRepository) UpdateReason(ctx context.Context, code string, reason string) error {
	if err := g.db.First(&entities.Appointment{}, code).Error; err != nil {
		return err
	}
	return g.db.Model(&entities.Appointment{}).Where("code = ?", code).Update("reason", reason).Error

}

// UpdateStatus implements AppointmentRepository.
func (g *gormAppointmentRepository) UpdateStatus(ctx context.Context, code string, status string) error {
	if err := g.db.First(&entities.Appointment{}, code).Error; err != nil {
		return err
	}
	return g.db.Model(&entities.Appointment{}).Where("code = ?", code).Update("status", status).Error
}

// UpdateType implements AppointmentRepository.
func (g *gormAppointmentRepository) UpdateType(ctx context.Context, code string, appointmentType string) error {
	if err := g.db.First(&entities.Appointment{}, code).Error; err != nil {
		return err
	}
	return g.db.Model(&entities.Appointment{}).Where("code = ?", code).Update("type", appointmentType).Error

}

func (g *gormAppointmentRepository) GenerateCode() (string, error) {
	var lastAppointment entities.Appointment
	if err := g.db.Order("id DESC").First(&lastAppointment).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return "AP-1", nil
		}
		return "", nil
	}

	newCode, err := utils.GenerateCode(lastAppointment.Code)
	if err != nil {
		return "", err
	}
	return newCode, nil
}

func (r *gormAppointmentRepository) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(ctx)
	})
}
