package repository

import (
	"warehouse/entity"
	"warehouse/model/response"
)

//SupplierRepository ...
type SupplierRepository interface {
	Insert(request *entity.Suppliers) (response entity.Suppliers)
	GetAll() (suppliers []response.SupplierResponse)
	GetById(id interface{}) (response entity.Suppliers)
	Update(request *entity.Suppliers) (response entity.Suppliers)
	DeleteById(id interface{})
}
