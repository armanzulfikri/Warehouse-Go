package request

//CreateRackRequest ...
type CreateRackRequest struct {
	WarehouseID  uint   `json:"warehouse_id`
	CategoryID   uint   `json:"category_id`
	RackCode     string `json:"rack_code"`
	RackCapacity int    `json:"rack_capacity"`
}
