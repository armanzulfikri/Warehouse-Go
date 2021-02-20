package repository

import (
	"warehouse/entity"
	"warehouse/model/response"
)

//WarehouseRepository ...
type WarehouseRepository interface {
	Insert(request *entity.Warehouses) (response entity.Warehouses)
	GetAll() (warehouse []response.WarehousesGetAllResponse)
	GetById(id interface{}) (response entity.Warehouses)
	Update(request *entity.Warehouses) (response entity.Warehouses)
	DeleteById(id interface{})
}
