package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"warehouse/entity"
	"warehouse/model/request"
	"warehouse/model/response"
	"warehouse/repository"
)

//NewUserService ...
func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceImpl{
		Repository: *userRepository,
	}
}

type userServiceImpl struct {
	Repository repository.UserRepository
}

//List
func (service userServiceImpl) List() (responses []response.UserResponse) {
	responses = service.Repository.GetAll()

	if len(responses) > 0 {
		fmt.Println(len(responses))
	} else {
		fmt.Println("FAILED")
	}

	return responses
}

//Create
func (service userServiceImpl) Create(request request.RegisterRequest) (response response.UserResponse) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	user := entity.Users{
		FullName:   request.FullName,
		Gender:     request.Gender,
		BirthDate:  request.BirthDate,
		BirthPlace: request.BirthPlace,
		Email:      request.FullName,
		Password:   string(hash),
		Role:       request.Role,
	}

	result := service.Repository.Insert(&user)

	if result.Email != "" {
		response.ID = result.ID
		response.Email = result.Email
		response.FullName = result.FullName
		response.Gender = result.Gender
		response.BirthDate = result.BirthDate
		response.BirthPlace = result.BirthPlace
		response.Email = result.FullName
		response.Role = result.Role
	}

	return
}

//Update
func (service userServiceImpl) Update(id interface{}, request request.RegisterRequest) (response response.UserResponse) {
	user := service.Repository.GetById(id)

	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	user.FullName = request.FullName
	user.Gender = request.Gender
	user.BirthDate = request.BirthDate
	user.BirthPlace = request.BirthPlace
	user.Email = request.FullName
	user.Password = string(hash)
	user.Role = request.Role

	service.Repository.Update(&user)

	if user.Email != "" {
		response.ID = user.ID
		response.Email = user.Email
		response.FullName = user.FullName
		response.Gender = user.Gender
		response.BirthDate = user.BirthDate
		response.BirthPlace = user.BirthPlace
		response.Email = user.FullName
		response.Role = user.Role
	}

	return
}

//GetById
func (service userServiceImpl) GetById(id interface{}) (response response.UserResponse) {
	result := service.Repository.GetById(id)

	if result.Email != "" {
		response.ID = result.ID
		response.Email = result.Email
		response.FullName = result.FullName
		response.Gender = result.Gender
		response.BirthDate = result.BirthDate
		response.BirthPlace = result.BirthPlace
		response.Email = result.FullName
		response.Role = result.Role
	}

	return
}

//DeleteById
func (service userServiceImpl) DeleteById(id interface{}) {
	service.Repository.DeleteById(id)
}
