package repository

import (
	"warehouse/entity"
	"warehouse/model/response"
)

//TransactionDetailRepository ...
type TransactionDetailRepository interface {
	Insert(request *entity.TransactionDetails) (response entity.TransactionDetails)
	GetAll() (resp []response.DetailResponse)
	GetByWarehouse(id interface{}) (resp []response.DetailResponse)
	GetByTransaction(id interface{}) (resp []response.DetailResponse)
	DeleteById(id interface{})
}
