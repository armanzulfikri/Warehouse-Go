package service

import (
	"warehouse/model/request"
	"warehouse/model/response"
)

//WarehouseService ...
type WarehouseService interface {
	List() (responses []response.WarehousesGetAllResponse)
	ListByProvince(provinceID interface{}) (responses []response.WarehousesGetAllResponse)
	Create(request request.CreateWarehouseRequest) (response response.WarehouseResponse)
	Update(id interface{}, request request.CreateWarehouseRequest) (response response.WarehouseResponse)
	GetById(id interface{}) (response response.WarehouseResponse)
	DeleteById(id interface{})
}
