package repository

import (
	"fmt"
	"warehouse/entity"
	"warehouse/model/response"

	"gorm.io/gorm"
)

//NewTransactionDetailRepository ...
func NewTransactionDetailRepository(database *gorm.DB) TransactionDetailRepository {
	return &transactionDetailRepositoryImpl{
		Database: database,
	}
}

type transactionDetailRepositoryImpl struct {
	Database *gorm.DB
}

func (repository transactionDetailRepositoryImpl) GetByWarehouse(id interface{}) (response []response.DetailResponse) {
	result := repository.Database.Model(&entity.TransactionDetails{}).Where("warehouses.id = ?", id).
		Select(`transaction_details.id, transaction_details.transaction_id, transaction_details.created_at, transaction_details.product_code,
				transaction_details.date_of_type, transaction_details.date, transaction_details.quantity, products.product_name,
				products.description, products.product_image_url, products.product_name, racks.rack_code`).
		Joins(`join transactions on transactions.id = transaction_details.transaction_id
				join products on transactions.product_id = products.id
				join racks on transactions.rack_id = racks.id
				join warehouses on racks.warehouse_id = warehouses.id`).Scan(&response)

	if result.RowsAffected > 0 {
		return
	} else {
		fmt.Println("Error ", result.Error)
	}

	return
}

func (repository transactionDetailRepositoryImpl) GetByTransaction(id interface{}) (response []response.DetailResponse) {
	result := repository.Database.Model(&entity.TransactionDetails{}).Where("transactions.id = ?", id).
		Select(`transaction_details.id, transaction_details.transaction_id, transaction_details.created_at, transaction_details.product_code,
			transaction_details.date_of_type, transaction_details.date, transaction_details.quantity, products.product_name,
			products.description, products.product_image_url, products.product_name, racks.rack_code`).
		Joins(`join transactions on transactions.id = transaction_details.transaction_id
			join products on transactions.product_id = products.id
			join racks on transactions.rack_id = racks.id
			join warehouses on racks.warehouse_id = warehouses.id`).Scan(&response)

	if result.RowsAffected > 0 {
		return
	} else {
		fmt.Println("Error ", result.Error)
	}

	return
}

func (repository transactionDetailRepositoryImpl) GetAll() (response []response.DetailResponse) {
	result := repository.Database.Model(&entity.TransactionDetails{}).Select(`transaction_details.id,
			transaction_details.transaction_id, transaction_details.created_at, transaction_details.product_code,
			transaction_details.date_of_type, transaction_details.date, transaction_details.quantity, products.product_name,
			products.description, products.product_image_url, products.product_name, racks.rack_code`).
		Joins(`join transactions on transactions.id = transaction_details.transaction_id
			join products on transactions.product_id = products.id
			join racks on transactions.rack_id = racks.id
			join warehouses on racks.warehouse_id = warehouses.id`).Scan(&response)

	if result.RowsAffected > 0 {
		return
	} else {
		fmt.Println("Error ", result.Error)
	}

	return
}

func (repository transactionDetailRepositoryImpl) Insert(request *entity.TransactionDetails) (response entity.TransactionDetails) {
	err := repository.Database.Create(&request).Scan(&response)

	if err.RowsAffected > 0 {
		return response
	} else {
		fmt.Println("error ", err.Error)
	}

	return entity.TransactionDetails{}
}

func (repository transactionDetailRepositoryImpl) DeleteById(id interface{}) {
	result := repository.Database.Delete(&entity.TransactionDetails{}, id)

	if result.RowsAffected > 0 {
		fmt.Println("Transaction Deleted")
	}

	fmt.Println("Fail deleted data ", result.Error)
}
