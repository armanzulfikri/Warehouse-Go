package repository

import (
	"fmt"
	"warehouse/entity"
	"warehouse/model/response"

	"gorm.io/gorm"
)

//NewDistrictRepository ...
func NewDistrictRepository(database *gorm.DB) DistrictRepository {
	return &districtRepositoryImpl{
		Database: database,
	}
}

type districtRepositoryImpl struct {
	Database *gorm.DB
}

func (repository districtRepositoryImpl) Insert(request *entity.Districts) (response entity.Districts) {
	err := repository.Database.Create(&request).Scan(&response)

	if err.RowsAffected > 0 {
		return response
	} else {
		fmt.Println("error ", err.Error)
	}

	return entity.Districts{}
}

func (repository districtRepositoryImpl) GetAll() (response []response.DistrictResponse) {
	result := repository.Database.Model(&entity.Districts{}).
		Select("districts.*").
		Scan(&response)

	if result.RowsAffected > 0 {

	} else {
		fmt.Println("Errors ", result.Error)
	}

	return
}

func (repository districtRepositoryImpl) GetById(id interface{}) (response entity.Districts) {
	result := repository.Database.First(&response, id)

	if result.RowsAffected > 0 {
		return
	} else {
		fmt.Println("Error ", result.Error)
	}

	return
}

func (repository districtRepositoryImpl) Update(request *entity.Districts) (response entity.Districts) {
	result := repository.Database.Save(&request).Scan(&response)
	if result.RowsAffected > 0 {
		return
	} else {
		fmt.Println("Failed to update date, error", result.Error)
	}

	return
}

func (repository districtRepositoryImpl) DeleteById(id interface{}) {
	result := repository.Database.Delete(&entity.Districts{}, id)

	if result.RowsAffected > 0 {
		fmt.Println("District Deleted")
	}

	fmt.Println("Fail deleted data ", result.Error)
}
