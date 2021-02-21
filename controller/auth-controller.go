package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse/model"
	"warehouse/model/request"
	"warehouse/service"
)

//AuthController ...
type AuthController struct {
	AuthService service.AuthService
}

//NewAuthController ...
func NewAuthController(authService *service.AuthService) AuthController {
	return AuthController{AuthService: *authService}
}

//Route ...
func (controller *AuthController) Route(router *gin.Engine) {
	router.POST("/auth/register", controller.Register)
	router.POST("/auth/login", controller.Login)
}

//Register ...
func (controller *AuthController) Register(context *gin.Context) {
	var request request.RegisterRequest
	context.Bind(&request)

	resp := controller.AuthService.Register(request)

	if resp.Email == "" {
		context.JSON(http.StatusOK,
			model.WebResponse{
				Code:   http.StatusBadRequest,
				Status: "Login Failed",
				Data:   nil,
			})
	} else {
		context.JSON(http.StatusOK,
			model.WebResponse{
				Code:   http.StatusOK,
				Status: "Register Success",
				Data:   resp,
			})
	}
}

//Login ...
func (controller *AuthController) Login(context *gin.Context) {
	var request request.LoginRequest

	context.Bind(&request)

	resp := controller.AuthService.Login(request)

	if resp.Token == "" {
		context.JSON(http.StatusOK,
			model.WebResponse{
				Code:   401,
				Status: "Login Failed",
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
