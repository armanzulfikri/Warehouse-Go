package entity

import (
	"time"

	"gorm.io/gorm"
)

//Suppliers struct
type Suppliers struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	SupplierName string         `json:"supplier_name"`
	Address      string         `json:"address"`
	Telepon      string         `json:"role"`
}
