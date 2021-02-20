package repository

import (
	"warehouse/entity"
	"warehouse/model/response"
)

//DistrictRepository ...
type DistrictRepository interface {
	Insert(request *entity.Districts) (response entity.Districts)
	GetAll() (district []response.DistrictResponse)
	GetById(id interface{}) (response entity.Districts)
	Update(request *entity.Districts) (response entity.Districts)
	DeleteById(id interface{})
}
