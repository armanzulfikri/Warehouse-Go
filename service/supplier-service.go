package service

import (
	"warehouse/model/request"
	"warehouse/model/response"
)

//SupplierService ...
type SupplierService interface {
	List() (responses []response.SupplierResponse)
	Create(request request.CreateSupplierRequest) (response response.SupplierResponse)
	Update(id interface{}, request request.CreateSupplierRequest) (response response.SupplierResponse)
	GetById(id interface{}) (response response.SupplierResponse)
	DeleteById(id interface{})
}
