package repository

import "warehouse/entity"

//ProductRepository ...
type ProductRepository interface {
	Insert(request entity.Products) (product entity.Products)
	GetAll() (products []entity.Products)
	GetById(id int) (product entity.Products)
	Update(request entity.Products) (product entity.Products)
	DeleteById(id int)
}