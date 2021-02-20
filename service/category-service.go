package service

import (
	"warehouse/model/request"
	"warehouse/model/response"
)

//CategoryService ...
type CategoryService interface {
	List() (responses []response.CategoryResponse)
	Create(request request.CreateCategoryRequest) (response response.CategoryResponse)
	Update(id interface{}, request request.CreateCategoryRequest) (response response.CategoryResponse)
	GetById(id interface{}) (response response.CategoryResponse)
	DeleteById(id interface{})
}
