package entity

import (
	"time"

	"gorm.io/gorm"
)

// TransactionDetails model / product details in warehouse
type TransactionDetails struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	TransactionID uint
	ProductCode   string
	DateOfType    string
	Date          time.Time
	Quantity      int
}
