package response

import "time"

//RackResponse ...
type RackResponse struct {
	ID           uint      `json:"id"`
	WarehouseID  uint      `json:"warehouse_id"`
	CategoryID   uint      `json:"category_id"`
	RackCode     string    `json:"rack_code"`
	RackCapacity int       `json:"rack_capacity"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}
