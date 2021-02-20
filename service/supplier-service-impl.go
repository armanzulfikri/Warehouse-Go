package service

import (
	"fmt"
	"warehouse/entity"
	"warehouse/model/request"
	"warehouse/model/response"
	"warehouse/repository"
)

//NewSupplierService ...
func NewSupplierService(supplierRepository *repository.SupplierRepository) SupplierService {
	return &supplierServiceImpl{
		Repository: *supplierRepository,
	}
}

type supplierServiceImpl struct {
	Repository repository.SupplierRepository
}

//List
func (service supplierServiceImpl) List() (responses []response.SupplierResponse) {
	responses = service.Repository.GetAll()

	if len(responses) > 0 {
		fmt.Println(len(responses))
	} else {
		fmt.Println("FAILED")
	}

	return responses
}

//Create
func (service supplierServiceImpl) Create(request request.CreateSupplierRequest) (response response.SupplierResponse) {
	supplier := entity.Suppliers{
		SupplierName: request.SupplierName,
		Address:      request.Address,
		Telepon:      request.Telepon,
	}

	result := service.Repository.Insert(&supplier)

	if result.SupplierName != "" {
		response.ID = result.ID
		response.SupplierName = result.SupplierName
		response.Address = result.Address
		response.Telepon = result.Telepon
	}

	return
}

//Update
func (service supplierServiceImpl) Update(WarehouseID interface{}, request request.CreateSupplierRequest) (response response.SupplierResponse) {
	supplier := service.Repository.GetById(WarehouseID)

	supplier.SupplierName = request.SupplierName
	supplier.Address = request.Address
	supplier.Telepon = request.Telepon

	service.Repository.Update(&supplier)

	if supplier.SupplierName != "" {
		response.ID = supplier.ID
		response.SupplierName = supplier.SupplierName
		response.Address = supplier.Address
		response.Telepon = supplier.Telepon
	}

	return
}

//GetById
func (service supplierServiceImpl) GetById(id interface{}) (response response.SupplierResponse) {
	result := service.Repository.GetById(id)

	if result.SupplierName != "" {
		response.ID = result.ID
		response.SupplierName = result.SupplierName
		response.Address = result.Address
		response.Telepon = result.Telepon

	}

	return
}

//DeleteById
func (service supplierServiceImpl) DeleteById(id interface{}) {
	service.Repository.DeleteById(id)
}
