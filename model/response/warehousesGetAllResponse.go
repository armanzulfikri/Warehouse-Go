package response

import (
	"time"
)

//WarehousesGetAllResponse ...
type WarehousesGetAllResponse struct {
	ID                 uint
	DistrictsID        uint   `json:"districts_id"`
	DistrictsName      string `json:"districts_name"`
	WarehousesName     string `json:"warehouses_name"`
	WarehousesCapacity int    `json:"warehouses_capacity"`
	WarehousesType     string `json:"warehouses_type"`
	WarehousesLocation string `json:"warehouses_location"`
	CreatedAt          time.Time
}
