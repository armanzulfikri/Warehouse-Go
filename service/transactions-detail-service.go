package service

import (
	"warehouse/model/request"
	"warehouse/model/response"
)

//TransactionDetailService ...
type TransactionDetailService interface {
	ListAll() (responses []response.DetailResponse)
	ListByWarehouse(id interface{}) (responses []response.DetailResponse)
	ListByTrans(id interface{}) (responses []response.DetailResponse)
	Create(request request.CreateDetailRequest) (response response.CreateDetailResponse)
	// GetById(id interface{}) (response response.TransactionResponse)
	DeleteById(id interface{})
}
