package request

//CreateProductRequest ...
type CreateProductRequest struct {
	UserID          uint   `json:"user_id"`
	CategoryID      uint   `json:"category_id"`
	SupplierID      uint   `json:"supplier_id"`
	ProductName     string `json:"product_name"`
	ProductImageURL string `json:"product_image_url"`
	Description     string `json:"description"`
}
