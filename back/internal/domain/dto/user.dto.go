package dto

type UserDTO struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	CI        string `json:"ci"`
	Birthdate string `json:"birthdate"`
	Role      string `json:"role"`
}
