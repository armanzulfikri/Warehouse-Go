package controller

import (
	"net/http"
	"warehouse/model"
	"warehouse/model/request"
	"warehouse/model/response"
	"warehouse/service"

	"github.com/gin-gonic/gin"
)

//WarehouseController ...
type WarehouseController struct {
	WarehouseService service.WarehouseService
}

//NewWarehouseController ...
func NewWarehouseController(provinceService *service.WarehouseService) WarehouseController {
	return WarehouseController{WarehouseService: *provinceService}
}

//Route ...
func (controller *WarehouseController) Route(group *gin.RouterGroup) {
	group.GET("/api/warehouse", controller.List)
	// group.GET("/api/warehouse/:province_id", controller.ListByProvince)
	group.GET("/api/warehouse/:id", controller.GetOne)
	group.POST("/api/warehouse", controller.Create)
	group.PUT("/api/warehouse/:id", controller.Update)
	group.DELETE("/api/warehouse/:id", controller.Delete)
}

//Create ...
func (controller *WarehouseController) Create(context *gin.Context) {
	var request request.CreateWarehouseRequest
	context.Bind(&request)

	resp := controller.WarehouseService.Create(request)

	if resp.WarehousesName == "" {
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
func (controller *WarehouseController) GetOne(context *gin.Context) {
	resp := controller.WarehouseService.GetById(context.Param("id"))

	if resp.WarehousesName == "" {
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
func (controller *WarehouseController) Update(context *gin.Context) {
	var request request.CreateWarehouseRequest

	context.Bind(&request)

	resp := controller.WarehouseService.Update(context.Param("id"), request)

	if resp.WarehousesName == "" {
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
func (controller *WarehouseController) List(context *gin.Context) {
	var resp []response.WarehousesGetAllResponse
	provID := context.Query("province_id")
	if provID == "" {
		resp = controller.WarehouseService.List()
	}
	if provID != "" {
		resp = controller.WarehouseService.ListByProvince(provID)
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

//ListByProvince ...
// func (controller *WarehouseController) ListByProvince(context *gin.Context) {
// 	resp := controller.WarehouseService.ListByProvince(context.Param(":province_id"))

// 	if len(resp) > 0 {
// 		context.JSON(http.StatusOK,
// 			model.WebResponse{
// 				Code:   http.StatusOK,
// 				Status: "OK",
// 				Data:   resp,
// 			})
// 	} else {
// 		context.JSON(http.StatusOK,
// 			model.WebResponse{
// 				Code:   http.StatusOK,
// 				Status: "OK",
// 				Data:   nil,
// 			})
// 	}
// }

//Delete ...
func (controller *WarehouseController) Delete(context *gin.Context) {

	controller.WarehouseService.DeleteById(context.Param("id"))

	context.JSON(http.StatusOK,
		model.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   nil,
		})
}
