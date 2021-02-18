package repository

import (
	"gorm.io/gorm"
	"warehouse/entity"
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
func (repository productRepositoryImpl) Insert(request entity.Products) (product entity.Products) {
	panic("implement me")
}

//GetAll ...
func (repository productRepositoryImpl) GetAll() (products []entity.Products) {
	panic("implement me")
}

//GetById ...
func (repository productRepositoryImpl) GetById(id int) (product entity.Products) {
	panic("implement me")
}

//Update
func (repository productRepositoryImpl) Update(product entity.Products) {
	panic("implement me")
}

//DeleteById
func (repository productRepositoryImpl) DeleteById(id int) {
	panic("implement me")
}
