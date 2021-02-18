package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse/model"
	"warehouse/model/request"
	"warehouse/service"
)

//ProductController ...
type ProductController struct {
	ProductService service.ProductService
}

//NewProductController ...
func NewProductController(productService *service.ProductService) ProductController {
	return ProductController{ProductService: *productService}
}

//Route ...
func (controller *ProductController) Route(group *gin.RouterGroup) {
	group.POST("/api/products", controller.Create)
	group.GET("/api/products", controller.List)
}

//Create ...
func (controller *ProductController) Create(context *gin.Context) {
	var request request.CreateProduct

	resp := controller.ProductService.Create(request)

	context.JSON(http.StatusOK,
		model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   resp,
	})
}

//List ...
func (controller *ProductController) List(context *gin.Context) {
	responses := controller.ProductService.List()
	context.JSON(
		http.StatusOK,
		model.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   responses,
		})
}
