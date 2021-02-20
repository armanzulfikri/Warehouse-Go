package service

import (
	"fmt"
	"warehouse/entity"
	"warehouse/model/request"
	"warehouse/model/response"
	"warehouse/repository"
)

//NewDistrictService ...
func NewDistrictService(districtRepository *repository.DistrictRepository) DistrictService {
	return &districtServiceImpl{
		Repository: *districtRepository,
	}
}

type districtServiceImpl struct {
	Repository repository.DistrictRepository
}

//List
func (service districtServiceImpl) List() (responses []response.DistrictResponse) {
	responses = service.Repository.GetAll()

	if len(responses) > 0 {
		fmt.Println(len(responses))
	} else {
		fmt.Println("FAILED")
	}

	return responses
}

//Create
func (service districtServiceImpl) Create(request request.CreateDistrictRequest) (response response.DistrictResponse) {
	district := entity.Districts{
		Name: request.Name,
	}

	result := service.Repository.Insert(&district)

	if result.Name != "" {
		response.ID = result.ID
		response.Name = result.Name
	}

	return
}

//Update
func (service districtServiceImpl) Update(id interface{}, request request.CreateDistrictRequest) (response response.DistrictResponse) {
	district := service.Repository.GetById(id)

	district.Name = request.Name

	service.Repository.Update(&district)

	if district.Name != "" {
		response.ID = district.ID
		response.Name = district.Name
	}

	return
}

//GetById
func (service districtServiceImpl) GetById(id interface{}) (response response.DistrictResponse) {
	result := service.Repository.GetById(id)

	if result.Name != "" {
		response.ID = result.ID
		response.Name = result.Name
	}

	return
}

//DeleteById
func (service districtServiceImpl) DeleteById(id interface{}) {
	service.Repository.DeleteById(id)
}
