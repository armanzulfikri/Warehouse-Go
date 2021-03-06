package response

import (
	"time"

	"gorm.io/gorm"
)

//WarehouseResponse ...
type WarehouseResponse struct {
	ID                 uint
	DistrictsID        uint
	WarehousesName     string `json:"warehouses_name"`
	WarehousesCapacity int    `json:"warehouses_capacity"`
	WarehousesType     string `json:"warehouses_type"`
	WarehousesLocation string `json:"warehouses_location"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}
