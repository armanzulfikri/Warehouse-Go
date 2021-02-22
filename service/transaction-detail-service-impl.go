package service

import (
	"fmt"
	"time"
	"warehouse/entity"
	"warehouse/model/request"
	"warehouse/model/response"
	"warehouse/repository"
)

//NewTransactionDetailService ...
func NewTransactionDetailService(transactionDetailRepository *repository.TransactionDetailRepository) TransactionDetailService {
	return &transactionDetailServiceImpl{
		Repository: *transactionDetailRepository,
	}
}

type transactionDetailServiceImpl struct {
	Repository repository.TransactionDetailRepository
}

//List
func (service transactionDetailServiceImpl) ListAll() (responses []response.DetailResponse) {
	responses = service.Repository.GetAll()

	if len(responses) > 0 {
		fmt.Println(len(responses))
	} else {
		fmt.Println("FAILED")
	}

	return responses
}

//List
func (service transactionDetailServiceImpl) ListByWarehouse(id interface{}) (responses []response.DetailResponse) {
	responses = service.Repository.GetByWarehouse(id)

	if len(responses) > 0 {
		fmt.Println(len(responses))
	} else {
		fmt.Println("FAILED")
	}

	return responses
}

//List
func (service transactionDetailServiceImpl) ListByTrans(transactionID interface{}) (responses []response.DetailResponse) {
	responses = service.Repository.GetByTransaction(transactionID)

	if len(responses) > 0 {
		fmt.Println(len(responses))
	} else {
		fmt.Println("FAILED")
	}

	return responses
}

//Create
func (service transactionDetailServiceImpl) Create(request request.CreateDetailRequest) (response response.CreateDetailResponse) {
	transaction := entity.TransactionDetails{
		TransactionID: request.TransactionID,
		ProductCode:   request.ProductCode,
		DateOfType:    request.DateOfType,
		Quantity:      request.Quantity,
		Date:          time.Now(),
	}

	result := service.Repository.Insert(&transaction)

	if result.Quantity != 0 {
		response.ID = result.ID
		response.TransactionID = result.TransactionID
		response.ProductCode = result.ProductCode
		response.DateOfType = result.DateOfType
		response.Quantity = result.Quantity
		response.Date = result.Date
	}

	return
}

//DeleteById
func (service transactionDetailServiceImpl) DeleteById(id interface{}) {
	service.Repository.DeleteById(id)
}
