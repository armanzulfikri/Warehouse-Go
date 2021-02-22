package service

import (
	"fmt"
	"strconv"
	"warehouse/entity"
	"warehouse/model/request"
	"warehouse/model/response"
	"warehouse/repository"

	"github.com/dgrijalva/jwt-go"
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
func (service productServiceImpl) Create(request request.CreateProductRequest, payload jwt.MapClaims) (response response.ProductResponse) {
	id := payload["id"]
	strID := fmt.Sprintf("%v", id)
	uintID, _ := strconv.ParseUint(strID, 10, 32)
	fmt.Println(uintID)
	product := entity.Products{
		UserID:          uint(uintID),
		CategoryID:      request.CategoryID,
		SupplierID:      request.SupplierID,
		ProductName:     request.ProductName,
		ProductImageURL: request.ProductImageURL,
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
