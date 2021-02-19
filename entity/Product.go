package entity

import (
	"time"

	"gorm.io/gorm"
)

//Product example ...
type Product struct {
	Id   int
	Name string
}

// Products true table
type Products struct {
	ID              uint `gorm:"primarykey"`
	UserID          uint
	CategoryID      uint
	SupplierID      uint
	ProductName     string
	ProductImageURL string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
