package dto

type LoginDTO struct {
	Email   string `json:"email" validate:"required,email"`
	Message string `json:"message"`
}
