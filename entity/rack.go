package entity

import "gorm.io/gorm"

//Racks Entity
type Racks struct {
	gorm.Model
	WarehousesID uint
	CategoryID   uint
	RackCode     string `json:"rack_code"`
	RackCapacity int    `json:"rack_capacity"`
}
