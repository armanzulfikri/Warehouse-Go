package controller

import (
	"net/http"
	"warehouse/model"
	"warehouse/model/request"
	"warehouse/service"

	"github.com/gin-gonic/gin"
)

//RackController ...
type RackController struct {
	RackService service.RackService
}

//NewRackController ...
func NewRackController(supplierService *service.RackService) RackController {
	return RackController{RackService: *supplierService}
}

//Route ...
func (controller *RackController) Route(group *gin.RouterGroup) {
	group.GET("/api/rack/", controller.GetWarehouses)
	group.GET("/api/rack", controller.List)
	group.GET("/api/rack/:id", controller.GetOne)
	group.POST("/api/rack", controller.Create)
	group.PUT("/api/rack/:id", controller.Update)
	group.DELETE("/api/rack/:id", controller.Delete)
}

//Create ...
func (controller *RackController) Create(context *gin.Context) {
	var request request.CreateRackRequest
	context.Bind(&request)

	resp := controller.RackService.Create(request)

	if resp.RackCode == "" {
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
func (controller *RackController) GetOne(context *gin.Context) {
	resp := controller.RackService.GetById(context.Param("id"))

	if resp.RackCode == "" {
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

//GetWarehouses ...
func (controller *RackController) GetWarehouses(context *gin.Context) {
	resp := controller.RackService.GetByWarehouses(context.Query("warehouse_id"))

	if resp.WarehouseID == 0 {
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
func (controller *RackController) Update(context *gin.Context) {
	var request request.CreateRackRequest

	context.Bind(&request)

	resp := controller.RackService.Update(context.Param("id"), request)

	if resp.RackCode == "" {
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
func (controller *RackController) List(context *gin.Context) {
	resp := controller.RackService.List()

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
func (controller *RackController) Delete(context *gin.Context) {

	controller.RackService.DeleteById(context.Param("id"))

	context.JSON(http.StatusOK,
		model.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   nil,
		})
}
