package response

import "time"

//ProductResponse ...
type ProductResponse struct {
	ID              uint
	UserID          uint      `json:"user_id"`
	CategoryID      uint      `json:"category_id"`
	SupplierID      uint      `json:"supplier_id"`
	ProductName     string    `json:"product_name"`
	ProductImageURL string    `json:"product_image"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}
