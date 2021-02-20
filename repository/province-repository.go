package repository

import (
	"warehouse/entity"
	"warehouse/model/response"
)

//ProvinceRepository ...
type ProvinceRepository interface {
	Insert(request *entity.Provinces) (response entity.Provinces)
	GetAll() (province []response.ProvinceResponse)
	GetById(id interface{}) (response entity.Provinces)
	Update(request *entity.Provinces) (response entity.Provinces)
	DeleteById(id interface{})
}
