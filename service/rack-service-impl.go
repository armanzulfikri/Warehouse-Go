package service

import (
	"fmt"
	"warehouse/entity"
	"warehouse/model/request"
	"warehouse/model/response"
	"warehouse/repository"
)

//NewRackService ...
func NewRackService(rackRepository *repository.RackRepository) RackService {
	return &rackServiceImpl{
		Repository: *rackRepository,
	}
}

type rackServiceImpl struct {
	Repository repository.RackRepository
}

//List
func (service rackServiceImpl) List() (responses []response.GetRackResponse) {
	responses = service.Repository.GetAll()

	if len(responses) > 0 {
		fmt.Println(len(responses))
	} else {
		fmt.Println("FAILED")
	}

	return responses
}

//Create
func (service rackServiceImpl) Create(request request.CreateRackRequest) (response response.RackResponse) {
	rack := entity.Racks{
		WarehouseID:  request.WarehouseID,
		CategoryID:   request.CategoryID,
		RackCode:     request.RackCode,
		RackCapacity: request.RackCapacity,
	}

	result := service.Repository.Insert(&rack)

	if result.RackCode != "" {
		response.ID = result.ID
		response.WarehouseID = result.WarehouseID
		response.CategoryID = result.CategoryID
		response.RackCode = result.RackCode
		response.RackCapacity = result.RackCapacity
	}

	return
}

//Update
func (service rackServiceImpl) Update(id interface{}, request request.CreateRackRequest) (response response.RackResponse) {
	rack := service.Repository.GetById(id)

	rack.WarehouseID = request.WarehouseID
	rack.CategoryID = request.CategoryID
	rack.RackCode = request.RackCode
	rack.RackCapacity = request.RackCapacity

	service.Repository.Update(&rack)

	if rack.RackCode != "" {
		response.ID = rack.ID
		response.WarehouseID = rack.WarehouseID
		response.CategoryID = rack.CategoryID
		response.RackCode = rack.RackCode
		response.RackCapacity = rack.RackCapacity
	}

	return
}

//GetById
func (service rackServiceImpl) GetById(id interface{}) (response response.RackResponse) {
	result := service.Repository.GetById(id)

	if result.RackCode != "" {
		response.ID = result.ID
		response.WarehouseID = result.WarehouseID
		response.CategoryID = result.CategoryID
		response.RackCode = result.RackCode
		response.RackCapacity = result.RackCapacity

	}

	return
}

//GetById
func (service rackServiceImpl) GetByWarehouses(warehouseid interface{}) (response response.RackResponse) {
	result := service.Repository.GetByWarehouses(warehouseid)

	if result.WarehouseID != 0 {
		response.ID = result.ID
		response.WarehouseID = result.WarehouseID
		response.CategoryID = result.CategoryID
		response.RackCode = result.RackCode
		response.RackCapacity = result.RackCapacity

	}

	return
}

//DeleteById
func (service rackServiceImpl) DeleteById(id interface{}) {
	service.Repository.DeleteById(id)
}
