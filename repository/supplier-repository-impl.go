package repository

import (
	"fmt"
	"warehouse/entity"
	"warehouse/model/response"

	"gorm.io/gorm"
)

//NewSupplierRepository ...
func NewSupplierRepository(database *gorm.DB) SupplierRepository {
	return &supplierRepositoryImpl{
		Database: database,
	}
}

type supplierRepositoryImpl struct {
	Database *gorm.DB
}

func (repository supplierRepositoryImpl) Insert(request *entity.Suppliers) (response entity.Suppliers) {
	err := repository.Database.Create(&request).Scan(&response)

	if err.RowsAffected > 0 {
		return response
	} else {
		fmt.Println("error ", err.Error)
	}

	return entity.Suppliers{}
}

func (repository supplierRepositoryImpl) GetAll() (response []response.SupplierResponse) {
	result := repository.Database.Model(&entity.Suppliers{}).
		Select("suppliers.*").
		Scan(&response)

	if result.RowsAffected > 0 {

	} else {
		fmt.Println("Errors ", result.Error)
	}

	return
}

func (repository supplierRepositoryImpl) GetById(id interface{}) (response entity.Suppliers) {
	result := repository.Database.First(&response, id)

	if result.RowsAffected > 0 {
		return
	} else {
		fmt.Println("Error ", result.Error)
	}

	return
}

func (repository supplierRepositoryImpl) Update(request *entity.Suppliers) (response entity.Suppliers) {
	result := repository.Database.Save(&request).Scan(&response)
	if result.RowsAffected > 0 {
		return
	} else {
		fmt.Println("Failed to update date, error", result.Error)
	}

	return
}

func (repository supplierRepositoryImpl) DeleteById(id interface{}) {
	result := repository.Database.Delete(&entity.Suppliers{}, id)

	if result.RowsAffected > 0 {
		fmt.Println("Supplier Deleted")
	}

	fmt.Println("Fail deleted data ", result.Error)
}
