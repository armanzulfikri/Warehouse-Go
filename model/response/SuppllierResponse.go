package response

import "time"

//SupplierResponse ...
type SupplierResponse struct {
	ID           uint
	SupplierName string    `json:"supplier_name"`
	Address      string    `json:"address"`
	Telepon      string    `json:"telepon"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}
