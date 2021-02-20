package repository

import (
	"warehouse/entity"
	"warehouse/model/response"
)

//CategoryRepository ...
type CategoryRepository interface {
	Insert(request *entity.Categories) (response entity.Categories)
	GetAll() (category []response.CategoryResponse)
	GetById(id interface{}) (response entity.Categories)
	Update(request *entity.Categories) (response entity.Categories)
	DeleteById(id interface{})
}
