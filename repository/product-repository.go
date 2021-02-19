package repository

import "warehouse/entity"

//ProductRepository ...
type ProductRepository interface {
	Insert(request entity.Product) (product entity.Product)
	GetAll() (products []entity.Product)
	GetById(id int) (product entity.Product)
	Update(request entity.Product) (product entity.Product)
	DeleteById(id int)
}
