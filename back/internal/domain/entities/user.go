package entities

import (
	"back/internal/utils"
	"regexp"

	"gorm.io/gorm"
)

// User represents a user entity
type User struct {
	gorm.Model
	Username     string        `json:"username" gorm:"unique;not null"`
	Birthdate    string        `json:"birthdate"`
	Password     string        `json:"password" gorm:"not null"`
	Email        string        `json:"email" gorm:"unique;not null"`
	Role         string        `json:"role" gorm:"not null"`
	CI           string        `json:"ci" gorm:"uniqueIndex;not null"`
	Appointments []Appointment `gorm:"foreignKey:PatientCI;references:CI"`
}

// BeforeCreate is a function to asign a role to a user before creating it
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Role = "user"
	return
}

// BeforeAdmin is a function to change the role of a user to admin
func (u *User) BeforeAdmin(tx *gorm.DB) (err error) {
	u.Role = "admin"
	return
}

// BeforeDelete is a function to change the role of a user to deleted
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	u.Role = "deleted"
	return
}

// ValidateAtCreate is a function to validate the user entity
func (u *User) ValidateAtCreate() error {

	if err := u.ValidateUsername(); err != nil {
		return err
	}

	if err := u.ValidatePassword(); err != nil {
		return err
	}

	if err := u.ValidateEmail(); err != nil {
		return err
	}
	if err := u.ValidateCI(); err != nil {
		return err
	}
	if err := u.ValidateBithdate(); err != nil {
		return err
	}

	return nil

}

// ValidateUsername is a function to validate the username of a user
func (u *User) ValidateUsername() error {
	if u.Username == "" {
		return &utils.ValidationError{Field: "username", Message: "Username is required"}
	}
	return nil
}

func (u *User) ValidatePassword() error {
	minLength := 6
	hasLower := false
	hasUpper := false
	hasDigit := false

	for _, char := range u.Password {
		if char >= 'a' && char <= 'z' {
			hasLower = true
		} else if char >= 'A' && char <= 'Z' {
			hasUpper = true
		} else if char >= '0' && char <= '9' {
			hasDigit = true
		}

		if hasLower && hasUpper && hasDigit {
			break
		}
	}
	if !hasLower || !hasUpper || !hasDigit || len(u.Password) < minLength {
		return &utils.ValidationError{Field: "password", Message: "Password must contain at least one lowercase letter, one uppercase letter, one digit and be at least 6 characters long"}
	}

	return nil
}

func (u *User) ValidateEmail() error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(u.Email) {
		return &utils.ValidationError{Field: "email", Message: "Invalid email"}
	}
	return nil
}

// ValidateCI is a function to validate the CI of a user
func (u *User) ValidateCI() error {
	if u.CI == "" {
		return &utils.ValidationError{Field: "ci", Message: "CI is required"}
	}
	if len(u.CI) < 6 {
		return &utils.ValidationError{Field: "ci", Message: "CI must have 6 digits"}
	}

	return nil
}
func (u *User) ValidateBithdate() error {
	if u.Birthdate == "" {
		return &utils.ValidationError{Field: "birthdate", Message: "Birthdate is required"}
	}
	return nil
}
