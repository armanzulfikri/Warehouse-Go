package response

import (
	"time"
)

//ProductGetAllResponse ...
type ProductGetAllResponse struct {
	ID           uint
	UserName     string `json:"user_name"`
	CategoryName string `json:category_name`
	SupplierName string `json:supplier_name`
	ProductName  string `json:product_name`
	Description  string `json:description`
	CreatedAt    time.Time
}
