package repository

import (
	"fmt"
	"gorm.io/gorm"
	"warehouse/entity"
	"warehouse/model/response"
)

//NewUserRepository ...
func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		Database: database,
	}
}

type userRepositoryImpl struct {
	Database *gorm.DB
}

func (repository userRepositoryImpl) GetByEmail(email string) (user entity.Users) {
	result := repository.Database.Where("email = ?", email).First(&user)

	if result.RowsAffected > 0 {
		return
	} else {
		fmt.Println("Error", result.Error)
	}

	return
}

func (repository userRepositoryImpl) Insert(request *entity.Users) (response entity.Users) {
	err := repository.Database.Create(&request).Scan(&response)

	if err.RowsAffected > 0 {
		return response
	} else {
		fmt.Println("error ", err.Error)
	}

	return entity.Users{}
}

func (repository userRepositoryImpl) GetAll() (response []response.UserResponse) {
	result := repository.Database.Model(&entity.Users{}).
		Select("users.*").
		Scan(&response)

	if result.RowsAffected > 0 {

	} else {
		fmt.Println("Errors ", result.Error)
	}

	return
}

func (repository userRepositoryImpl) GetById(id interface{}) (response entity.Users) {
	result := repository.Database.First(&response, id)

	if result.RowsAffected > 0 {
		return
	} else {
		fmt.Println("Error ", result.Error)
	}

	return
}

func (repository userRepositoryImpl) Update(request *entity.Users) (response entity.Users) {
	result := repository.Database.Save(&request).Scan(&response)
	if result.RowsAffected > 0 {
		return
	} else {
		fmt.Println("Failed to update date, error", result.Error)
	}

	return
}

func (repository userRepositoryImpl) DeleteById(id interface{}) {
	result := repository.Database.Delete(&entity.Users{}, id)

	if result.RowsAffected > 0 {
		fmt.Println("User Deleted")
	}

	fmt.Println("Fail deleted data ", result.Error)
}
