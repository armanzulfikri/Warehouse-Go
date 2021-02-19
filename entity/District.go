package entity

import (
	"time"

	"gorm.io/gorm"
)

// Districts model
type Districts struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Name       string
	ProvinceID uint
}
