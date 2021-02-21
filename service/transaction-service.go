package service

import (
	"warehouse/model/request"
	"warehouse/model/response"
)

//TransactionService ...
type TransactionService interface {
	List(warehouseID interface{}) (responses []response.TransactionResponse)
	Create(request request.CreateTransactionRequest) (response response.CreateTransactionResponse)
	// GetById(id interface{}) (response response.TransactionResponse)
	DeleteById(id interface{})
}
