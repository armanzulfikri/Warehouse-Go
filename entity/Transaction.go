package entity

import (
	"time"

	"gorm.io/gorm"
)

// Transactions model / product in warehouse
type Transactions struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	ProductID    uint
	UserID       uint
	RackID       uint
	ProductStock int
}
