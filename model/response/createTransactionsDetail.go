package response

import (
	"time"

	"gorm.io/gorm"
)

// CreateDetailResponse model
type CreateDetailResponse struct {
	ID            uint           `json:"id"`
	TransactionID uint           `json:"transaction_id"`
	ProductCode   string         `json:"product_code"`
	DateOfType    string         `json:"date_of_type"`
	Date          time.Time      `json:"date"`
	Quantity      int            `json:"quantity"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}
