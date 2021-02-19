package repository

import (
	"warehouse/entity"

	"gorm.io/gorm"
)

//NewProductRepository
func NewProductRepository(database *gorm.DB) ProductRepository {
	return &productRepositoryImpl{
		Database: database,
	}
}

type productRepositoryImpl struct {
	Database *gorm.DB
}

//Insert ...
func (repository productRepositoryImpl) Insert(request entity.Product) (product entity.Product) {
	panic("implement me")
}

//GetAll ...
func (repository productRepositoryImpl) GetAll() (products []entity.Product) {
	panic("implement me")
}

//GetById ...
func (repository productRepositoryImpl) GetById(id int) (product entity.Product) {
	repository.Database.First(&product, id)

	return product
}

//Update
func (repository productRepositoryImpl) Update(request entity.Product) (product entity.Product) {
	panic("implement me")
}

//DeleteById
func (repository productRepositoryImpl) DeleteById(id int) {
	panic("implement me")
}
