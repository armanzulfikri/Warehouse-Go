package controller

import (
	"net/http"
	"warehouse/model"
	"warehouse/service"

	"github.com/gin-gonic/gin"
)

//StatisticController ...
type StatisticController struct {
	StatisticService service.StatisticService
}

//NewStatisticController ...
func NewStatisticController(statisticService *service.StatisticService) StatisticController {
	return StatisticController{StatisticService: *statisticService}
}

//Route ...
func (controller *StatisticController) Route(group *gin.RouterGroup) {
	// group.GET("/api/transaction/:id", controller.GetOne)
	group.GET("/api/statistic", controller.List)
}

// List ...
func (controller *StatisticController) List(context *gin.Context) {
	resp := controller.StatisticService.List()

	if resp.Product > 0 {
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
