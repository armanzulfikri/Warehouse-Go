package service

import (
	"fmt"
	"warehouse/entity"
	"warehouse/model/request"
	"warehouse/model/response"
	"warehouse/repository"
)

//NewWarehouseService ...
func NewWarehouseService(warehouseRepository *repository.WarehouseRepository) WarehouseService {
	return &warehouseServiceImpl{
		Repository: *warehouseRepository,
	}
}

type warehouseServiceImpl struct {
	Repository repository.WarehouseRepository
}

//List
func (service warehouseServiceImpl) List() (responses []response.WarehousesGetAllResponse) {
	responses = service.Repository.GetAll()

	if len(responses) > 0 {
		fmt.Println(len(responses))
	} else {
		fmt.Println("FAILED")
	}

	return responses
}

//Create
func (service warehouseServiceImpl) Create(request request.CreateWarehouseRequest) (response response.WarehouseResponse) {
	warehouse := entity.Warehouses{
		DistrictsID:        request.DistrictsID,
		WarehousesName:     request.WarehousesName,
		WarehousesCapacity: request.WarehousesCapacity,
		WarehousesType:     request.WarehousesType,
		WarehousesLocation: request.WarehousesLocation,
	}

	result := service.Repository.Insert(&warehouse)

	if result.WarehousesName != "" {
		response.ID = result.ID
		response.DistrictsID = result.DistrictsID
		response.WarehousesName = result.WarehousesName
		response.WarehousesType = result.WarehousesType
		response.WarehousesCapacity = result.WarehousesCapacity
		response.WarehousesLocation = result.WarehousesLocation
	}

	return
}

//Update
func (service warehouseServiceImpl) Update(id interface{}, request request.CreateWarehouseRequest) (response response.WarehouseResponse) {
	warehouse := service.Repository.GetById(id)

	warehouse.DistrictsID = request.DistrictsID
	warehouse.WarehousesName = request.WarehousesName
	warehouse.WarehousesType = request.WarehousesType
	warehouse.WarehousesCapacity = request.WarehousesCapacity
	warehouse.WarehousesLocation = request.WarehousesLocation

	service.Repository.Update(&warehouse)

	if warehouse.WarehousesName != "" {
		response.ID = warehouse.ID
		response.DistrictsID = warehouse.DistrictsID
		response.WarehousesName = warehouse.WarehousesName
		response.WarehousesType = warehouse.WarehousesType
		response.WarehousesCapacity = warehouse.WarehousesCapacity
		response.WarehousesLocation = warehouse.WarehousesLocation
	}

	return
}

//GetById
func (service warehouseServiceImpl) GetById(id interface{}) (response response.WarehouseResponse) {
	result := service.Repository.GetById(id)

	if result.WarehousesName != "" {
		response.ID = result.ID
		response.DistrictsID = result.DistrictsID
		response.WarehousesName = result.WarehousesName
		response.WarehousesType = result.WarehousesType
		response.WarehousesCapacity = result.WarehousesCapacity
		response.WarehousesLocation = result.WarehousesLocation
	}

	return
}

//DeleteById
func (service warehouseServiceImpl) DeleteById(id interface{}) {
	service.Repository.DeleteById(id)
}
