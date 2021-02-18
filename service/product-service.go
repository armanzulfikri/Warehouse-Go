package service

import (
	"warehouse/model/request"
	"warehouse/model/response"
)

//ProductService ...
type ProductService interface {
	Create(request request.CreateProduct) (response response.CreateProduct)
	List() (responses []response.GetProduct)
}
