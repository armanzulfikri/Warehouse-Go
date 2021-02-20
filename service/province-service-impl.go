package service

import (
	"fmt"
	"warehouse/entity"
	"warehouse/model/request"
	"warehouse/model/response"
	"warehouse/repository"
)

//NewProvinceService ...
func NewProvinceService(provinceRepository *repository.ProvinceRepository) ProvinceService {
	return &provinceServiceImpl{
		Repository: *provinceRepository,
	}
}

type provinceServiceImpl struct {
	Repository repository.ProvinceRepository
}

//List
func (service provinceServiceImpl) List() (responses []response.ProvinceResponse) {
	responses = service.Repository.GetAll()

	if len(responses) > 0 {
		fmt.Println(len(responses))
	} else {
		fmt.Println("FAILED")
	}

	return responses
}

//Create
func (service provinceServiceImpl) Create(request request.CreateProvinceRequest) (response response.ProvinceResponse) {
	province := entity.Provinces{
		Name: request.Name,
	}

	result := service.Repository.Insert(&province)

	if result.Name != "" {
		response.ID = result.ID
		response.Name = result.Name
	}

	return
}

//Update
func (service provinceServiceImpl) Update(id interface{}, request request.CreateProvinceRequest) (response response.ProvinceResponse) {
	province := service.Repository.GetById(id)

	province.Name = request.Name

	service.Repository.Update(&province)

	if province.Name != "" {
		response.ID = province.ID
		response.Name = province.Name
	}

	return
}

//GetById
func (service provinceServiceImpl) GetById(id interface{}) (response response.ProvinceResponse) {
	result := service.Repository.GetById(id)

	if result.Name != "" {
		response.ID = result.ID
		response.Name = result.Name
	}

	return
}

//DeleteById
func (service provinceServiceImpl) DeleteById(id interface{}) {
	service.Repository.DeleteById(id)
}
