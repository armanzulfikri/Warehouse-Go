package web

import (
	"warehouse/config"
	"warehouse/controller"
	"warehouse/repository"
	"warehouse/service"

	"github.com/gin-gonic/gin"
)

//Route
func Route(router *gin.Engine) *gin.Engine {
	database := *config.GetConnection()

	/*
	 * Please Comment repo, service, and controller if code error
	 */

	// Setup Repository
	productRepository := repository.NewProductRepository(&database)
	userRepository := repository.NewUserRepository(&database)
	provinceRepository := repository.NewProvinceRepository(&database)
	districtRepository := repository.NewDistrictRepository(&database)
	categoryRepository := repository.NewCategoryRepository(&database)
	warehouseRepository := repository.NewWarehouseRepository(&database)

	// Setup Service
	productService := service.NewProductService(&productRepository)
	authService := service.NewAuthService(&userRepository)
	userService := service.NewUserService(&userRepository)
	provinceService := service.NewProvinceService(&provinceRepository)
	districtService := service.NewDistrictService(&districtRepository)
	categoryService := service.NewCategoryService(&categoryRepository)
	warehouseService := service.NewWarehouseService(&warehouseRepository)
	// Setup Controller
	productController := controller.NewProductController(&productService)
	authController := controller.NewAuthController(&authService)
	userController := controller.NewUserController(&userService)
	provinceController := controller.NewProvinceController(&provinceService)
	districtController := controller.NewDistrictController(&districtService)
	categoryController := controller.NewCategoryController(&categoryService)
	warehouseController := controller.NewWarehouseController(&warehouseService)
	// For Auth And Register
	authController.Route(router)

	group := *router.Group("/")

	group.Use(Middleware)
	{
		productController.Route(&group)
		provinceController.Route(&group)
		userController.Route(&group)
		districtController.Route(&group)
		categoryController.Route(&group)
		warehouseController.Route(&group)

	}

	return router
}
