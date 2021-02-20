package controller

import (
	"net/http"
	"warehouse/model"
	"warehouse/model/request"
	"warehouse/service"

	"github.com/gin-gonic/gin"
)

//SupplierController ...
type SupplierController struct {
	SupplierService service.SupplierService
}

//NewSupplierController ...
func NewSupplierController(supplierService *service.SupplierService) SupplierController {
	return SupplierController{SupplierService: *supplierService}
}

//Route ...
func (controller *SupplierController) Route(group *gin.RouterGroup) {
	group.GET("/api/supplier", controller.List)
	group.GET("/api/supplier/:id", controller.GetOne)
	group.POST("/api/supplier", controller.Create)
	group.PUT("/api/supplier/:id", controller.Update)
	group.DELETE("/api/supplier/:id", controller.Delete)
}

//Create ...
func (controller *SupplierController) Create(context *gin.Context) {
	var request request.CreateSupplierRequest
	context.Bind(&request)

	resp := controller.SupplierService.Create(request)

	if resp.SupplierName == "" {
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
func (controller *SupplierController) GetOne(context *gin.Context) {
	resp := controller.SupplierService.GetById(context.Param("id"))

	if resp.SupplierName == "" {
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
func (controller *SupplierController) Update(context *gin.Context) {
	var request request.CreateSupplierRequest

	context.Bind(&request)

	resp := controller.SupplierService.Update(context.Param("id"), request)

	if resp.SupplierName == "" {
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
func (controller *SupplierController) List(context *gin.Context) {
	resp := controller.SupplierService.List()

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
func (controller *SupplierController) Delete(context *gin.Context) {

	controller.SupplierService.DeleteById(context.Param("id"))

	context.JSON(http.StatusOK,
		model.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   nil,
		})
}
