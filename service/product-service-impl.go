package service

import (
	"warehouse/entity"
	"warehouse/model/request"
	"warehouse/model/response"
	"warehouse/repository"
)

func NewProductService(productRepository *repository.ProductRepository) ProductService {
	return &productServiceImpl{
		ProductRepository: *productRepository,
	}
}

type productServiceImpl struct {
	ProductRepository repository.ProductRepository
}

func (service productServiceImpl) Create(request request.CreateProduct) (resp response.CreateProduct) {
	//Tambah Validasi disini

	product := entity.Products{
		Id:       request.ID,
		Name:     request.Name,
	}

	service.ProductRepository.Insert(product)

	resp = response.CreateProduct{
		ID: product.Id,
		Name: product.Name,
	}

	return resp
}

func (service productServiceImpl) List() (responses []response.GetProduct) {
	return responses
}

