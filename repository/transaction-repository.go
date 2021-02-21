package repository

import (
	"warehouse/entity"
	"warehouse/model/response"
)

//TransactionRepository ...
type TransactionRepository interface {
	Insert(request *entity.Transactions) (response entity.Transactions)
	GetAll() (transaction []response.TransactionResponse)
	GetById(id interface{}) (response entity.Transactions)
	Update(request *entity.Transactions) (response entity.Transactions)
	DeleteById(id interface{})
}
