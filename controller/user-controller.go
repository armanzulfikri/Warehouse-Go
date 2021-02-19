package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse/model"
	"warehouse/model/request"
	"warehouse/service"
)

//UserController ...
type UserController struct {
	UserService service.UserService
}

//NewUserController ...
func NewUserController(userService *service.UserService) UserController {
	return UserController{UserService: *userService}
}

//Route ...
func (controller *UserController) Route(group *gin.RouterGroup) {
	group.GET("/api/users", controller.List)
	group.GET("/api/user/:id", controller.GetOne)
	group.POST("/api/user", controller.Create)
	group.PUT("/api/user/:id", controller.Update)
	group.DELETE("/api/user/:id", controller.Delete)
}

//Create ...
func (controller *UserController) Create(context *gin.Context) {
	var request request.RegisterRequest
	context.Bind(&request)

	resp := controller.UserService.Create(request)

	if resp.Email == "" {
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
func (controller *UserController) GetOne(context *gin.Context) {
	resp := controller.UserService.GetById(context.Param("id"))

	if resp.Email == "" {
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
func (controller *UserController) Update(context *gin.Context) {
	var request request.RegisterRequest

	context.Bind(&request)

	resp := controller.UserService.Update(context.Param("id"), request)

	if resp.Email == "" {
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
func (controller *UserController) List(context *gin.Context) {
	resp := controller.UserService.List()

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
func (controller *UserController) Delete(context *gin.Context) {

	controller.UserService.DeleteById(context.Param("id"))

	context.JSON(http.StatusOK,
		model.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   nil,
		})
}
