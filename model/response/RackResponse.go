package response

import "time"

//RackResponse ...
type RackResponse struct {
	ID           uint
	WarehouseID  uint      `json:"warehouses_id"`
	CategoryID   uint      `json:"categories_id"`
	RackCode     string    `json:"rack_code"`
	RackCapacity int       `json:"rack_capacity"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}
