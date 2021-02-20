package request

//CreateWarehouseRequest ...
type CreateWarehouseRequest struct {
	DistrictsID        uint
	WarehousesName     string `json:"warehouses_name"`
	WarehousesCapacity int    `json:"warehouses_capacity"`
	WarehousesType     string `json:"warehouses_type"`
	WarehousesLocation string `json:"warehouses_location"`
}
