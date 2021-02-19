package repository

import (
	"warehouse/entity"
	"warehouse/model/response"
)

//UserRepository ...
type UserRepository interface {
	Insert(request *entity.Users) (response entity.Users)
	GetAll() (users []response.UserResponse)
	GetByEmail(email string) (user entity.Users)
	GetById(id interface{}) (response entity.Users)
	Update(request *entity.Users) (response entity.Users)
	DeleteById(id interface{})
}
