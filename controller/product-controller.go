package controller

import (
	"net/http"
	"warehouse/model"
	"warehouse/model/request"
	"warehouse/model/response"
	"warehouse/service"

	"github.com/gin-gonic/gin"
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
	group.GET("/api/product", controller.List)
	group.GET("/api/product/:id", controller.GetOne)
	group.POST("/api/product", controller.Create)
	group.PUT("/api/product/:id", controller.Update)
	group.DELETE("/api/product/:id", controller.Delete)
}

//Create ...
func (controller *ProductController) Create(context *gin.Context) {
	var request request.CreateProductRequest
	context.Bind(&request)

	resp := controller.ProductService.Create(request)

	if resp.ProductName == "" {
		context.JSON(http.StatusOK,
			model.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "Create Failed",
				Data:   nil,
			})
	} else {
		context.JSON(http.StatusOK,
			model.WebResponse{
				Code:   http.StatusOK,
				Status: "Create Success",
				Data:   resp,
			})
	}
}

//GetOne ...
func (controller *ProductController) GetOne(context *gin.Context) {
	resp := controller.ProductService.GetById(context.Param("id"))

	if resp.ProductName == "" {
		context.JSON(http.StatusOK,
			model.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "Select Failed",
				Data:   nil,
			})
	} else {
		context.JSON(http.StatusOK,
			model.WebResponse{
				Code:   http.StatusOK,
				Status: "Select Success",
				Data:   resp,
			})
	}
}

//Update ...
func (controller *ProductController) Update(context *gin.Context) {
	var request request.CreateProductRequest

	context.Bind(&request)

	resp := controller.ProductService.Update(context.Param("id"), request)

	if resp.ProductName == "" {
		context.JSON(http.StatusOK,
			model.WebResponse{
				Code:   401,
				Status: "Update User Failed",
				Data:   nil,
			})
	} else {
		context.JSON(http.StatusOK,
			model.WebResponse{
				Code:   200,
				Status: "Login Success",
				Data:   resp,
			})
	}
}

//List ...
func (controller *ProductController) List(context *gin.Context) {
	var resp []response.ProductGetAllResponse
	suppID := context.Query("supplier_id")
	if suppID == "" {
		resp = controller.ProductService.List()
	}
	if suppID != "" {
		resp = controller.ProductService.ListBySupplier(suppID)
	}

	if len(resp) > 0 {
		context.JSON(http.StatusOK,
			model.WebResponse{
				Code:   http.StatusOK,
				Status: "OK",
				Data:   resp,
			})
	} else {
		context.JSON(http.StatusOK,
			model.WebResponse{
				Code:   http.StatusOK,
				Status: "OK",
				Data:   nil,
			})
	}
}

//Delete ...
func (controller *ProductController) Delete(context *gin.Context) {

	controller.ProductService.DeleteById(context.Param("id"))

	context.JSON(http.StatusOK,
		model.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   nil,
		})
}
