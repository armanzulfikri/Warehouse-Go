package entity

import (
	"time"

	"gorm.io/gorm"
)

//Users Entity
type Users struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	FullName   string         `json:"full_name"`
	Gender     string         `json:"gender"`
	BirthDate  string         `json:"birth_date"`
	BirthPlace string         `json:"birth_place"`
	Role       string         `json:"role"`
	Email      string         `json:"email"`
	Password   string         `json:"password"`
}
