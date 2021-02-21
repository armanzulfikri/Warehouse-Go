package repository

import (
	"fmt"
	"warehouse/entity"
	"warehouse/model/response"

	"gorm.io/gorm"
)

//NewTransactionRepository ...
func NewTransactionRepository(database *gorm.DB) TransactionRepository {
	return &transactionRepositoryImpl{
		Database: database,
	}
}

type transactionRepositoryImpl struct {
	Database *gorm.DB
}

func (repository transactionRepositoryImpl) GetByWarehouse(warehouseID interface{}) (response []response.TransactionResponse) {
	result := repository.Database.Model(&entity.Transactions{}).Where("warehouses.id = ?", warehouseID).
		Select("transactions.id, transactions.product_stock, products.product_name, products.description, products.product_image_url, products.product_name, racks.rack_code").
		Joins(`join products on transactions.product_id = products.id
	join racks on transactions.rack_id = racks.id
	join warehouses on racks.warehouse_id = warehouses.id`).Scan(&response)

	if result.RowsAffected > 0 {
		return
	} else {
		fmt.Println("Error ", result.Error)
	}

	return
}

// func (repository transactionRepositoryImpl) GetById(id interface{}) (response response.TransactionResponse) {
// 	result := repository.Database.Model(&entity.Transactions{}).Where("id = ?", id).
// 		Select("transactions.id, transactions.product_stock, products.product_name, products.description, products.product_image_url").
// 		Joins(`join products on transactions.product_id = products.id`).Scan(&response)

// 	if result.RowsAffected > 0 {
// 		return
// 	} else {
// 		fmt.Println("Error ", result.Error)
// 	}

// 	return
// }

func (repository transactionRepositoryImpl) Insert(request *entity.Transactions) (response entity.Transactions) {
	err := repository.Database.Create(&request).Scan(&response)

	if err.RowsAffected > 0 {
		return response
	} else {
		fmt.Println("error ", err.Error)
	}

	return entity.Transactions{}
}

func (repository transactionRepositoryImpl) DeleteById(id interface{}) {
	result := repository.Database.Delete(&entity.Transactions{}, id)

	if result.RowsAffected > 0 {
		fmt.Println("Transaction Deleted")
	}

	fmt.Println("Fail deleted data ", result.Error)
}
