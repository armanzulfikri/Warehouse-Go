package service

import (
	"fmt"
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

//List
func (service productServiceImpl) List() (responses []response.ProductGetAllResponse) {
	responses = service.Repository.GetAll()

	if len(responses) > 0 {
		fmt.Println(len(responses))
	} else {
		fmt.Println("FAILED")
	}

	return responses
}

//List
func (service productServiceImpl) ListBySupplier(id interface{}) (responses []response.ProductGetAllResponse) {
	responses = service.Repository.GetBySupplier(id)

	if len(responses) > 0 {
		fmt.Println(len(responses))
	} else {
		fmt.Println("FAILED")
	}

	return responses
}

//Create
func (service productServiceImpl) Create(request request.CreateProductRequest) (response response.ProductResponse) {
	product := entity.Products{
		UserID:          request.UserID,
		CategoryID:      request.CategoryID,
		SupplierID:      request.SupplierID,
		ProductName:     request.ProductName,
		ProductImageURL: "https://images.unsplash.com/photo-1494961104209-3c223057bd26?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=1292&q=80",
		Description:     request.Description,
	}

	result := service.Repository.Insert(&product)

	if result.ProductName != "" {
		response.ID = result.ID
		response.UserID = result.UserID
		response.CategoryID = result.CategoryID
		response.SupplierID = result.SupplierID
		response.ProductName = result.ProductName
		response.ProductImageURL = result.ProductImageURL
		response.Description = result.Description
	}

	return
}

//Update
func (service productServiceImpl) Update(id interface{}, request request.CreateProductRequest) (response response.ProductResponse) {
	product := service.Repository.GetById(id)

	product.UserID = request.UserID
	product.CategoryID = request.CategoryID
	product.SupplierID = request.SupplierID
	product.ProductName = request.ProductName
	product.ProductImageURL = request.ProductImageURL

	service.Repository.Update(&product)

	if product.ProductName != "" {
		response.ID = product.ID
		response.UserID = product.UserID
		response.CategoryID = product.CategoryID
		response.SupplierID = product.SupplierID
		response.ProductName = product.ProductName
		response.ProductImageURL = product.ProductImageURL

	}

	return
}

//GetById
func (service productServiceImpl) GetById(id interface{}) (response response.ProductResponse) {
	result := service.Repository.GetById(id)

	if result.ProductName != "" {
		response.ID = result.ID
		response.UserID = result.UserID
		response.CategoryID = result.CategoryID
		response.SupplierID = result.SupplierID
		response.ProductName = result.ProductName
		response.ProductImageURL = result.ProductImageURL

	}

	return
}

//DeleteById
func (service productServiceImpl) DeleteById(id interface{}) {
	service.Repository.DeleteById(id)
}
