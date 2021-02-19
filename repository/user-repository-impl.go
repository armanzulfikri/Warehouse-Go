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

func (repository userRepositoryImpl) GetByEmail(email string) (user entity.Users) {
	repository.Database.Where("email = ?", email).First(&user)

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

func (repository userRepositoryImpl) GetById(id int) (user entity.Users) {
	panic("implement me")
}

func (repository userRepositoryImpl) Update(request entity.Users) (response entity.Users) {
	panic("implement me")
}

func (repository userRepositoryImpl) DeleteById(id int) {
	panic("implement me")
}
