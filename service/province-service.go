package service

import (
	"warehouse/model/request"
	"warehouse/model/response"
)

//ProvinceService ...
type ProvinceService interface {
	List() (responses []response.ProvinceResponse)
	Create(request request.CreateProvinceRequest) (response response.ProvinceResponse)
	Update(id interface{}, request request.CreateProvinceRequest) (response response.ProvinceResponse)
	GetById(id interface{}) (response response.ProvinceResponse)
	DeleteById(id interface{})
}
