package entity

import (
	"time"

	"gorm.io/gorm"
)

//Racks Entity
type Racks struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	WarehouseID  uint           `json:"warehouse_id"`
	CategoryID   uint           `json:"category_id"`
	RackCode     string         `json:"rack_code"`
	RackCapacity int            `json:"rack_capacity"`
}
