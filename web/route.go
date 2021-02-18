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
	database := config.GetConnection()
	// Setup Repository
	productRepository := repository.NewProductRepository(database)

	// Setup Service
	productService := service.NewProductService(&productRepository)

	// Setup Controller
	productController := controller.NewProductController(&productService)

	group := *router.Group("/")

	group.Use(Middleware)
	{
		productController.Route(&group)
	}

	return router
}
