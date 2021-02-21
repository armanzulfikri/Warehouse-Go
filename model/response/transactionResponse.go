package response

// TransactionResponse model
type TransactionResponse struct {
	ID              uint   `json:"id"`
	ProductStock    int    `json:"product_stock"`
	Description     string `json:"description"`
	ProductImageURL string `json:"product_image_url"`
}
