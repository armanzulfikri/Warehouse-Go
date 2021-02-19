package service

import (
	"warehouse/entity"
	"warehouse/model/request"
	"warehouse/model/response"
	"warehouse/repository"
)

//NewProductService ...
func NewProductService(productRepository *repository.ProductRepository) ProductService {
	return &productServiceImpl{
		Repository: *productRepository,
	}
}

type productServiceImpl struct {
	Repository repository.ProductRepository
}

//Create ...
func (service productServiceImpl) Create(request request.CreateProduct) (resp response.CreateProduct) {
	//Tambah Validasi disini

	product := entity.Products{
		Id:       request.ID,
		Name:     request.Name,
	}

	service.Repository.Insert(product)

	resp = response.CreateProduct{
		ID: product.Id,
		Name: product.Name,
	}

	return resp
}

//List ...
func (service productServiceImpl) List() (responses []response.GetProduct) {
	return responses
}

