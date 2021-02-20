package controller

import (
	"net/http"
	"warehouse/model"
	"warehouse/model/request"
	"warehouse/service"

	"github.com/gin-gonic/gin"
)

//DistrictController ...
type DistrictController struct {
	DistrictService service.DistrictService
}

//NewDistrictController ...
func NewDistrictController(districtService *service.DistrictService) DistrictController {
	return DistrictController{DistrictService: *districtService}
}

//Route ...
func (controller *DistrictController) Route(group *gin.RouterGroup) {
	group.GET("/api/district", controller.List)
	group.GET("/api/district/:id", controller.GetOne)
	group.POST("/api/district", controller.Create)
	group.PUT("/api/district/:id", controller.Update)
	group.DELETE("/api/district/:id", controller.Delete)
}

//Create ...
func (controller *DistrictController) Create(context *gin.Context) {
	var request request.CreateDistrictRequest
	context.Bind(&request)

	resp := controller.DistrictService.Create(request)

	if resp.Name == "" {
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
func (controller *DistrictController) GetOne(context *gin.Context) {
	resp := controller.DistrictService.GetById(context.Param("id"))

	if resp.Name == "" {
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
func (controller *DistrictController) Update(context *gin.Context) {
	var request request.CreateDistrictRequest

	context.Bind(&request)

	resp := controller.DistrictService.Update(context.Param("id"), request)

	if resp.Name == "" {
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
func (controller *DistrictController) List(context *gin.Context) {
	resp := controller.DistrictService.List()

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
func (controller *DistrictController) Delete(context *gin.Context) {

	controller.DistrictService.DeleteById(context.Param("id"))

	context.JSON(http.StatusOK,
		model.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   nil,
		})
}
