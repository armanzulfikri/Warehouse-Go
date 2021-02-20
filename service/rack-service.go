package service

import (
	"warehouse/model/request"
	"warehouse/model/response"
)

//RackService ...
type RackService interface {
	List() (responses []response.GetRackResponse)
	Create(request request.CreateRackRequest) (response response.RackResponse)
	Update(id interface{}, request request.CreateRackRequest) (response response.RackResponse)
	GetById(id interface{}) (response response.RackResponse)
	GetByWarehouses(warehouseid interface{}) (response response.RackResponse)
	DeleteById(id interface{})
}
