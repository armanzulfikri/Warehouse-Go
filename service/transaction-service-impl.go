package service

import (
	"fmt"
	"warehouse/entity"
	"warehouse/model/request"
	"warehouse/model/response"
	"warehouse/repository"
)

//NewTransactionService ...
func NewTransactionService(transactionRepository *repository.TransactionRepository) TransactionService {
	return &transactionServiceImpl{
		Repository: *transactionRepository,
	}
}

type transactionServiceImpl struct {
	Repository repository.TransactionRepository
}

//List
func (service transactionServiceImpl) List(warehouseID interface{}) (responses []response.TransactionResponse) {
	responses = service.Repository.GetByWarehouse(warehouseID)

	if len(responses) > 0 {
		fmt.Println(len(responses))
	} else {
		fmt.Println("FAILED")
	}

	return responses
}

//Create
func (service transactionServiceImpl) Create(request request.CreateTransactionRequest) (response response.CreateTransactionResponse) {
	transaction := entity.Transactions{
		ProductID:    request.ProductID,
		RackID:       request.RackID,
		ProductStock: 0,
	}

	result := service.Repository.Insert(&transaction)

	if result.ProductID != 0 {
		response.ID = result.ID
		response.ProductID = result.ProductID
		response.RackID = result.RackID
		response.ProductStock = result.ProductStock
	}

	return
}

// //GetById
// func (service transactionServiceImpl) GetById(id interface{}) (response response.TransactionResponse) {
// 	result := service.Repository.GetById(id)

// 	if result.Description != "" {
// 		response.ID = result.ID
// 		response.Description = result.Description
// 		response.ProductImageURL = result.ProductImageURL
// 		response.ProductStock = result.ProductStock

// 	}

// 	return
// }

//DeleteById
func (service transactionServiceImpl) DeleteById(id interface{}) {
	service.Repository.DeleteById(id)
}
