package repository

import (
	"warehouse/entity"
	"warehouse/model/response"
)

//RackRepository ...
type RackRepository interface {
	Insert(request *entity.Racks) (response entity.Racks)
	GetAll() (racks []response.GetRackResponse)
	GetById(id interface{}) (response entity.Racks)
	GetByWarehouses(warehouseid interface{}) (response entity.Racks)
	Update(request *entity.Racks) (response entity.Racks)
	DeleteById(id interface{})
}
