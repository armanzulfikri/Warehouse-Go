package response

import "time"

// DetailResponse model
type DetailResponse struct {
	ID              uint      `json:"id"`
	TransactionID   uint      `json:"transaction_id"`
	ProductName     string    `json:"product_name"`
	ProductImageURL string    `json:"product_image_url"`
	ProductCode     string    `json:"product_code"`
	DateOfType      string    `json:"date_of_type"`
	Quantity        int       `json:"quantity"`
	Date            time.Time `json:"date"`
	CreatedAt       time.Time `json:"created_at"`
}
