package repository

import (
	"warehouse/entity"
	"warehouse/model/response"
)

//TransactionRepository ...
type TransactionRepository interface {
	Insert(request *entity.Transactions) (response entity.Transactions)
	GetByWarehouse(warehouseID interface{}) (transaction []response.TransactionResponse)
	// GetById(id interface{}) (response response.TransactionResponse)
	DeleteById(id interface{})
}
