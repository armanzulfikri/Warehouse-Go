package service

import (
	"warehouse/model/request"
	"warehouse/model/response"
)

//ProductService ...
type ProductService interface {
	List() (responses []response.ProductGetAllResponse)
	ListBySupplier(id interface{}) (responses []response.ProductGetAllResponse)
	Create(request request.CreateProductRequest) (response response.ProductResponse)
	Update(id interface{}, request request.CreateProductRequest) (response response.ProductResponse)
	GetById(id interface{}) (response response.ProductResponse)
	DeleteById(id interface{})
}
