package service

import (
	"warehouse/model/request"
	"warehouse/model/response"
)

//DistrictService ...
type DistrictService interface {
	List() (responses []response.DistrictResponse)
	Create(request request.CreateDistrictRequest) (response response.DistrictResponse)
	Update(id interface{}, request request.CreateDistrictRequest) (response response.DistrictResponse)
	GetById(id interface{}) (response response.DistrictResponse)
	DeleteById(id interface{})
}
