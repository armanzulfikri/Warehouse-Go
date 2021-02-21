package request

// CreateDetailRequest model
type CreateDetailRequest struct {
	TransactionID uint   `json:"transaction_id"`
	ProductCode   string `json:"product_code"`
	DateOfType    string `json:"date_of_type"`
	Quantity      int    `json:"quantity"`
}