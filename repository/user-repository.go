package repository

import (
	"warehouse/entity"
)

//UserRepository ...
type UserRepository interface {
	Insert(request *entity.Users) (response entity.Users)
	GetAll() (users []entity.Users)
	GetByEmail(email string) (user entity.Users)
	GetById(id interface{}) (response entity.Users)
	Update(request *entity.Users) (response entity.Users)
	DeleteById(id interface{})
}
