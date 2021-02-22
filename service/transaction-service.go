package service

import (
	"warehouse/model/request"
	"warehouse/model/response"

	"github.com/dgrijalva/jwt-go"
)

//TransactionService ...
type TransactionService interface {
	List(warehouseID interface{}) (responses []response.TransactionResponse)
	Create(request request.CreateTransactionRequest, payload jwt.MapClaims) (response response.CreateTransactionResponse)
	// GetById(id interface{}) (response response.TransactionResponse)
	DeleteById(id interface{})
}
