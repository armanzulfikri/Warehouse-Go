package response

// TransactionResponse model
type TransactionResponse struct {
	ID              uint   `json:"id"`
	ProductName     string `json:"product_name"`
	RackCode        string `json:"rack_code"`
	ProductStock    int    `json:"product_stock"`
	Description     string `json:"description"`
	ProductImageURL string `json:"product_image_url"`
}
