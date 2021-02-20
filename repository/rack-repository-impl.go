package repository

import (
	"fmt"
	"warehouse/entity"
	"warehouse/model/response"

	"gorm.io/gorm"
)

//NewRackRepository ...
func NewRackRepository(database *gorm.DB) RackRepository {
	return &rackRepositoryImpl{
		Database: database,
	}
}

type rackRepositoryImpl struct {
	Database *gorm.DB
}

func (repository rackRepositoryImpl) Insert(request *entity.Racks) (response entity.Racks) {
	err := repository.Database.Create(&request).Scan(&response)

	if err.RowsAffected > 0 {
		return response
	} else {
		fmt.Println("error ", err.Error)
	}

	return entity.Racks{}
}

func (repository rackRepositoryImpl) GetAll() (response []response.GetRackResponse) {
	result := repository.Database.Model(&entity.Racks{}).
		Select("racks.id, warehouses.warehouses_name,categories.category_name,racks.rack_code,racks.rack_capacity").
		Joins("join warehouses on racks.warehouse_id = warehouses.id").
		Joins("join categories on racks.category_id = categories.id").
		Scan(&response)

	if result.RowsAffected > 0 {

	} else {
		fmt.Println("Errors ", result.Error)
	}

	return
}

func (repository rackRepositoryImpl) GetById(id interface{}) (response entity.Racks) {
	result := repository.Database.First(&response, id)

	if result.RowsAffected > 0 {
		return
	} else {
		fmt.Println("Error ", result.Error)
	}

	return
}

func (repository rackRepositoryImpl) GetByWarehouses(warehouseid interface{}) (response entity.Racks) {
	result := repository.Database.Where("warehouse_id = ?", warehouseid).Find(&response)

	if result.RowsAffected > 0 {
		return
	} else {
		fmt.Println("Error ", result.Error)
	}

	return
}
func (repository rackRepositoryImpl) Update(request *entity.Racks) (response entity.Racks) {
	result := repository.Database.Save(&request).Scan(&response)
	if result.RowsAffected > 0 {
		return
	} else {
		fmt.Println("Failed to update date, error", result.Error)
	}

	return
}

func (repository rackRepositoryImpl) DeleteById(id interface{}) {
	result := repository.Database.Delete(&entity.Racks{}, id)

	if result.RowsAffected > 0 {
		fmt.Println("Rack Deleted")
	}

	fmt.Println("Fail deleted data ", result.Error)
}
