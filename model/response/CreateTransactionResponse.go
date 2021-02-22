package response

import (
	"time"

	"gorm.io/gorm"
)

// CreateTransactionResponse model
type CreateTransactionResponse struct {
	ID           uint           `json:"id"`
	UserID       uint           `json:"user_id"`
	ProductID    uint           `json:"product_id"`
	RackID       uint           `json:"rack_id"`
	ProductStock int            `json:"product_stock"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}
