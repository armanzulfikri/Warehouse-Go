package repository

import (
	"fmt"
	"warehouse/entity"
	"warehouse/model/response"

	"gorm.io/gorm"
)

//NewProductRepository ...
func NewProductRepository(database *gorm.DB) ProductRepository {
	return &productRepositoryImpl{
		Database: database,
	}
}

type productRepositoryImpl struct {
	Database *gorm.DB
}

func (repository productRepositoryImpl) Insert(request *entity.Products) (response entity.Products) {
	err := repository.Database.Create(&request).Scan(&response)

	if err.RowsAffected > 0 {
		return response
	} else {
		fmt.Println("error ", err.Error)
	}

	return entity.Products{}
}

func (repository productRepositoryImpl) GetAll() (response []response.ProductGetAllResponse) {
	result := repository.Database.Model(&entity.Products{}).
		Select("products.id, users.full_name as input_by, categories.category_name,suppliers.supplier_name,products.product_name, products.product_image_url,products.description,products.created_at").
		Joins("join users on products.user_id = users.id").
		Joins("join categories on products.category_id = categories.id").
		Joins("join suppliers on products.supplier_id = suppliers.id").
		Scan(&response)

	if result.RowsAffected > 0 {

	} else {
		fmt.Println("Errors ", result.Error)
	}

	return
}

func (repository productRepositoryImpl) GetById(id interface{}) (response entity.Products) {
	result := repository.Database.First(&response, id)

	if result.RowsAffected > 0 {
		return
	} else {
		fmt.Println("Error ", result.Error)
	}

	return
}

func (repository productRepositoryImpl) Update(request *entity.Products) (response entity.Products) {
	result := repository.Database.Save(&request).Scan(&response)
	if result.RowsAffected > 0 {
		return
	} else {
		fmt.Println("Failed to update date, error", result.Error)
	}

	return
}

func (repository productRepositoryImpl) DeleteById(id interface{}) {
	result := repository.Database.Delete(&entity.Products{}, id)

	if result.RowsAffected > 0 {
		fmt.Println("Product Deleted")
	}

	fmt.Println("Fail deleted data ", result.Error)
}
