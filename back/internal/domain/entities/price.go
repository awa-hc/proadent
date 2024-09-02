package entities

import "gorm.io/gorm"

type Price struct {
	gorm.Model
	
	Price       float64 `json:"price" gorm:"not null"`
	Type        string  `json:"type" gorm:"not null"`
	Title       string  `json:"title" gorm:"not null"`
	Description string  `json:"description" gorm:"not null"`
	Active      bool    `json:"active" gorm:"not null"`
}
