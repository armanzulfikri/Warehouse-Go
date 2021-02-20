package repository

import (
	"fmt"
	"warehouse/entity"
	"warehouse/model/response"

	"gorm.io/gorm"
)

//NewProvinceRepository ...
func NewProvinceRepository(database *gorm.DB) ProvinceRepository {
	return &provinceRepositoryImpl{
		Database: database,
	}
}

type provinceRepositoryImpl struct {
	Database *gorm.DB
}

func (repository provinceRepositoryImpl) Insert(request *entity.Provinces) (response entity.Provinces) {
	err := repository.Database.Create(&request).Scan(&response)

	if err.RowsAffected > 0 {
		return response
	} else {
		fmt.Println("error ", err.Error)
	}

	return entity.Provinces{}
}

func (repository provinceRepositoryImpl) GetAll() (response []response.ProvinceResponse) {
	result := repository.Database.Model(&entity.Provinces{}).
		Select("provinces.*").
		Scan(&response)

	if result.RowsAffected > 0 {

	} else {
		fmt.Println("Errors ", result.Error)
	}

	return
}

func (repository provinceRepositoryImpl) GetById(id interface{}) (response entity.Provinces) {
	result := repository.Database.First(&response, id)

	if result.RowsAffected > 0 {
		return
	} else {
		fmt.Println("Error ", result.Error)
	}

	return
}

func (repository provinceRepositoryImpl) Update(request *entity.Provinces) (response entity.Provinces) {
	result := repository.Database.Save(&request).Scan(&response)
	if result.RowsAffected > 0 {
		return
	} else {
		fmt.Println("Failed to update date, error", result.Error)
	}

	return
}

func (repository provinceRepositoryImpl) DeleteById(id interface{}) {
	result := repository.Database.Delete(&entity.Provinces{}, id)

	if result.RowsAffected > 0 {
		fmt.Println("Province Deleted")
	}

	fmt.Println("Fail deleted data ", result.Error)
}
