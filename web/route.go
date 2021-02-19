package web

import (
	"github.com/gin-gonic/gin"
	"warehouse/config"
	"warehouse/controller"
	"warehouse/repository"
	"warehouse/service"
)

//Route
func Route(router *gin.Engine) *gin.Engine  {
	database := *config.GetConnection()

	/*
	 * Please Comment repo, service, and controller if code error
	 */


	// Setup Repository
	productRepository := repository.NewProductRepository(&database)
	userRepository := repository.NewUserRepository(&database)

	// Setup Service
	productService := service.NewProductService(&productRepository)
	authService := service.NewAuthService(&userRepository)
	userService := service.NewUserService(&userRepository)

	// Setup Controller
	productController := controller.NewProductController(&productService)
	authController := controller.NewAuthController(&authService)
	userController := controller.NewUserController(&userService)

	group := *router.Group("/")

	group.Use(Middleware)
	{
		productController.Route(&group)
		authController.Route(&group)
		userController.Route(&group)
	}

	return router
}
