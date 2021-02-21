package request

// CreateTransactionRequest model
type CreateTransactionRequest struct {
	ProductID uint `json:"product_id"`
	RackID    uint `json:"rack_id"`
}
