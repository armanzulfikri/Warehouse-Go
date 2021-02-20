package controller

import (
	"net/http"
	"warehouse/model"
	"warehouse/model/request"
	"warehouse/service"

	"github.com/gin-gonic/gin"
)

//ProvinceController ...
type ProvinceController struct {
	ProvinceService service.ProvinceService
}

//NewProvinceController ...
func NewProvinceController(provinceService *service.ProvinceService) ProvinceController {
	return ProvinceController{ProvinceService: *provinceService}
}

//Route ...
func (controller *ProvinceController) Route(group *gin.RouterGroup) {
	group.GET("/api/province", controller.List)
	group.GET("/api/province/:id", controller.GetOne)
	group.POST("/api/province", controller.Create)
	group.PUT("/api/province/:id", controller.Update)
	group.DELETE("/api/province/:id", controller.Delete)
}

//Create ...
func (controller *ProvinceController) Create(context *gin.Context) {
	var request request.CreateProvinceRequest
	context.Bind(&request)

	resp := controller.ProvinceService.Create(request)

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
func (controller *ProvinceController) GetOne(context *gin.Context) {
	resp := controller.ProvinceService.GetById(context.Param("id"))

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
func (controller *ProvinceController) Update(context *gin.Context) {
	var request request.CreateProvinceRequest

	context.Bind(&request)

	resp := controller.ProvinceService.Update(context.Param("id"), request)

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
func (controller *ProvinceController) List(context *gin.Context) {
	resp := controller.ProvinceService.List()

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
func (controller *ProvinceController) Delete(context *gin.Context) {

	controller.ProvinceService.DeleteById(context.Param("id"))

	context.JSON(http.StatusOK,
		model.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   nil,
		})
}
