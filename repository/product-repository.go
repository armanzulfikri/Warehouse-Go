package repository

import (
	"warehouse/entity"
	"warehouse/model/response"
)

//ProductRepository ...
type ProductRepository interface {
	Insert(request *entity.Products) (response entity.Products)
	GetAll() (product []response.ProductGetAllResponse)
	GetById(id interface{}) (response entity.Products)
	Update(request *entity.Products) (response entity.Products)
	DeleteById(id interface{})
}
