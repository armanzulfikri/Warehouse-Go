package service

import (
	"warehouse/model/request"
	"warehouse/model/response"
)

//UserService ...
type UserService interface {
	List() (responses []response.UserResponse)
	Create(request request.RegisterRequest) (response response.UserResponse)
	Update(id interface{}, request request.RegisterRequest) (response response.UserResponse)
	GetById(id interface{}) (response response.UserResponse)
	DeleteById(id interface{})
}
