package repository

import (
	"fmt"
	"warehouse/entity"
	"warehouse/model/response"

	"gorm.io/gorm"
)

//NewCategoryRepository ...
func NewCategoryRepository(database *gorm.DB) CategoryRepository {
	return &categoryRepositoryImpl{
		Database: database,
	}
}

type categoryRepositoryImpl struct {
	Database *gorm.DB
}

func (repository categoryRepositoryImpl) Insert(request *entity.Categories) (response entity.Categories) {
	err := repository.Database.Create(&request).Scan(&response)

	if err.RowsAffected > 0 {
		return response
	} else {
		fmt.Println("error ", err.Error)
	}

	return entity.Categories{}
}

func (repository categoryRepositoryImpl) GetAll() (response []response.CategoryResponse) {
	result := repository.Database.Model(&entity.Categories{}).
		Select("categories.*").
		Scan(&response)

	if result.RowsAffected > 0 {

	} else {
		fmt.Println("Errors ", result.Error)
	}

	return
}

func (repository categoryRepositoryImpl) GetById(id interface{}) (response entity.Categories) {
	result := repository.Database.First(&response, id)

	if result.RowsAffected > 0 {
		return
	} else {
		fmt.Println("Error ", result.Error)
	}

	return
}

func (repository categoryRepositoryImpl) Update(request *entity.Categories) (response entity.Categories) {
	result := repository.Database.Save(&request).Scan(&response)
	if result.RowsAffected > 0 {
		return
	} else {
		fmt.Println("Failed to update date, error", result.Error)
	}

	return
}

func (repository categoryRepositoryImpl) DeleteById(id interface{}) {
	result := repository.Database.Delete(&entity.Categories{}, id)

	if result.RowsAffected > 0 {
		fmt.Println("Category Deleted")
	}

	fmt.Println("Fail deleted data ", result.Error)
}
