package repository

import (
	"fmt"
	"warehouse/entity"
	"warehouse/model/response"

	"gorm.io/gorm"
)

//NewWarehouseRepository ...
func NewWarehouseRepository(database *gorm.DB) WarehouseRepository {
	return &warehouseRepositoryImpl{
		Database: database,
	}
}

type warehouseRepositoryImpl struct {
	Database *gorm.DB
}

func (repository warehouseRepositoryImpl) Insert(request *entity.Warehouses) (response entity.Warehouses) {
	err := repository.Database.Create(&request).Scan(&response)

	if err.RowsAffected > 0 {
		return response
	} else {
		fmt.Println("error ", err.Error)
	}

	return entity.Warehouses{}
}

func (repository warehouseRepositoryImpl) GetAll() (response []response.WarehousesGetAllResponse) {
	result := repository.Database.Model(&entity.Warehouses{}).
		Select(`warehouses.id, warehouses.districts_id as district_id, districts.name as district_name,
		warehouses.created_at, warehouses.updated_at, warehouses.deleted_at, warehouses.warehouses_name as warehouse_name, warehouses.warehouses_capacity as warehouse_capacity,
		warehouses.warehouses_type as warehouse_type, warehouses.warehouses_location as warehouse_location`).
		Joins("join districts on warehouses.districts_id = districts.id").
		Scan(&response)

	if result.RowsAffected > 0 {

	} else {
		fmt.Println("Errors ", result.Error)
	}

	return
}

func (repository warehouseRepositoryImpl) GetById(id interface{}) (response entity.Warehouses) {
	result := repository.Database.First(&response, id)

	if result.RowsAffected > 0 {
		return
	} else {
		fmt.Println("Error ", result.Error)
	}

	return
}

func (repository warehouseRepositoryImpl) Update(request *entity.Warehouses) (response entity.Warehouses) {
	result := repository.Database.Save(&request).Scan(&response)
	if result.RowsAffected > 0 {
		return
	} else {
		fmt.Println("Failed to update date, error", result.Error)
	}

	return
}

func (repository warehouseRepositoryImpl) DeleteById(id interface{}) {
	result := repository.Database.Delete(&entity.Warehouses{}, id)

	if result.RowsAffected > 0 {
		fmt.Println("Warehouse Deleted")
	}

	fmt.Println("Fail deleted data ", result.Error)
}
