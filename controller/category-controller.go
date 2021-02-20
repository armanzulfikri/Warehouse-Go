package controller

import (
	"net/http"
	"warehouse/model"
	"warehouse/model/request"
	"warehouse/service"

	"github.com/gin-gonic/gin"
)

//CategoryController ...
type CategoryController struct {
	CategoryService service.CategoryService
}

//NewCategoryController ...
func NewCategoryController(categoryService *service.CategoryService) CategoryController {
	return CategoryController{CategoryService: *categoryService}
}

//Route ...
func (controller *CategoryController) Route(group *gin.RouterGroup) {
	group.GET("/api/category", controller.List)
	group.GET("/api/category/:id", controller.GetOne)
	group.POST("/api/category", controller.Create)
	group.PUT("/api/category/:id", controller.Update)
	group.DELETE("/api/category/:id", controller.Delete)
}

//Create ...
func (controller *CategoryController) Create(context *gin.Context) {
	var request request.CreateCategoryRequest
	context.Bind(&request)

	resp := controller.CategoryService.Create(request)

	if resp.CategoryName == "" {
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
func (controller *CategoryController) GetOne(context *gin.Context) {
	resp := controller.CategoryService.GetById(context.Param("id"))

	if resp.CategoryName == "" {
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
func (controller *CategoryController) Update(context *gin.Context) {
	var request request.CreateCategoryRequest

	context.Bind(&request)

	resp := controller.CategoryService.Update(context.Param("id"), request)

	if resp.CategoryName == "" {
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
func (controller *CategoryController) List(context *gin.Context) {
	resp := controller.CategoryService.List()

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
func (controller *CategoryController) Delete(context *gin.Context) {

	controller.CategoryService.DeleteById(context.Param("id"))

	context.JSON(http.StatusOK,
		model.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   nil,
		})
}
