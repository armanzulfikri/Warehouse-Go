package response

import (
	"time"

	"gorm.io/gorm"
)

//WarehousesGetAllResponse ...
type WarehousesGetAllResponse struct {
	ID                uint           `json:"id"`
	DistrictID        uint           `json:"district_id"`
	DistrictName      string         `json:"district_name"`
	WarehouseName     string         `json:"warehouse_name"`
	WarehouseCapacity int            `json:"warehouse_capacity"`
	WarehouseType     string         `json:"warehouse_type"`
	WarehouseLocation string         `json:"warehouse_location"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"udpated_at"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at"`
}
