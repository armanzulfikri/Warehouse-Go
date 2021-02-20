package service

import (
	"fmt"
	"warehouse/entity"
	"warehouse/model/request"
	"warehouse/model/response"
	"warehouse/repository"
)

//NewCategoryService ...
func NewCategoryService(categoryRepository *repository.CategoryRepository) CategoryService {
	return &categoryServiceImpl{
		Repository: *categoryRepository,
	}
}

type categoryServiceImpl struct {
	Repository repository.CategoryRepository
}

//List
func (service categoryServiceImpl) List() (responses []response.CategoryResponse) {
	responses = service.Repository.GetAll()

	if len(responses) > 0 {
		fmt.Println(len(responses))
	} else {
		fmt.Println("FAILED")
	}

	return responses
}

//Create
func (service categoryServiceImpl) Create(request request.CreateCategoryRequest) (response response.CategoryResponse) {
	category := entity.Categories{
		CategoryName: request.CategoryName,
	}

	result := service.Repository.Insert(&category)

	if result.CategoryName != "" {
		response.ID = result.ID
		response.CategoryName = result.CategoryName
	}

	return
}

//Update
func (service categoryServiceImpl) Update(id interface{}, request request.CreateCategoryRequest) (response response.CategoryResponse) {
	category := service.Repository.GetById(id)

	category.CategoryName = request.CategoryName

	service.Repository.Update(&category)

	if category.CategoryName != "" {
		response.ID = category.ID
		response.CategoryName = category.CategoryName
	}

	return
}

//GetById
func (service categoryServiceImpl) GetById(id interface{}) (response response.CategoryResponse) {
	result := service.Repository.GetById(id)

	if result.CategoryName != "" {
		response.ID = result.ID
		response.CategoryName = result.CategoryName
	}

	return
}

//DeleteById
func (service categoryServiceImpl) DeleteById(id interface{}) {
	service.Repository.DeleteById(id)
}
