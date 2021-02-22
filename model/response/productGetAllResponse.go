package response

import (
	"time"
)

//ProductGetAllResponse ...
type ProductGetAllResponse struct {
	ID              uint   `json:"id"`
	InputBy         string `json:"input_by"`
	CategoryName    string `json:"category_name"`
	SupplierName    string `json:"supplier_name"`
	ProductName     string `json:"product_name"`
	Description     string `json:"description"`
	ProductImageURL string `json:"product_image_url"`
	CreatedAt       time.Time
}
