package entity

import "gorm.io/gorm"

//Warehouses entity
type Warehouses struct {
	gorm.Model
	DistrictsID        uint
	WarehousesName     string `json:"warehouses_name"`
	WarehousesCapacity int    `json:"warehouses_capacity"`
	WarehousesType     string `json:"warehouses_type"`
	WarehousesLocation string `json:"warehouses_location"`
}
