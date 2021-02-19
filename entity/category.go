package entity

import (
	"time"

	"gorm.io/gorm"
)

//Categories Entity
type Categories struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	CategoryName string         `json:"category_name"`
}
