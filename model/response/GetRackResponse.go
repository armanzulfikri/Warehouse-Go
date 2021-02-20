package response

import "time"

//GetRackResponse ...
type GetRackResponse struct {
	ID             uint
	WarehousesName string    `json:"warehouses_name"`
	CategoryName   string    `json:"category_name"`
	RackCode       string    `json:"rack_code"`
	RackCapacity   int       `json:"rack_capacity"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}
