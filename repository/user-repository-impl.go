package repository

import (
	"fmt"
	"gorm.io/gorm"
	"warehouse/entity"
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

func (repository userRepositoryImpl) GetByEmail(email string) (user *entity.Users) {
	result := repository.Database.Where("email = ?", email)

	if result.RowsAffected > 0 {
		result.First(&user)
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

func (repository userRepositoryImpl) GetAll() (users []entity.Users) {
	panic("implement me")
}

func (repository userRepositoryImpl) GetById(id interface{}) (user *entity.Users) {
	result := repository.Database.First(&user, id)
	if result.RowsAffected > 0 {
		result.Scan(&user)
	} else {
		fmt.Println("Failed to update date", result.Error)
	}

	return
}

func (repository userRepositoryImpl) Update(request *entity.Users) (response entity.Users) {
	result := repository.Database.Save(&request)
	if result.RowsAffected > 0 {
		result.Scan(&response)
	} else {
		fmt.Println("Failed to update date", result.Error)
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
