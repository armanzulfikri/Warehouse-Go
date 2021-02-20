package service

import (
	"warehouse/model/request"
	"warehouse/model/response"
)

//ProductService ...
type ProductService interface {
	List() (responses []response.ProductResponse)
	Create(request request.CreateProductRequest) (response response.ProductResponse)
	Update(id interface{}, request request.CreateProductRequest) (response response.ProductResponse)
	GetById(id interface{}) (response response.ProductResponse)
	DeleteById(id interface{})
}
