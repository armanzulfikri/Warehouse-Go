package controller

import (
	"log"
	"net/http"
	"warehouse/model"
	"warehouse/model/request"
	"warehouse/service"

	"github.com/gin-gonic/gin"
)

//TransactionController ...
type TransactionController struct {
	TransactionService service.TransactionService
}

//NewTransactionController ...
func NewTransactionController(transactionService *service.TransactionService) TransactionController {
	return TransactionController{TransactionService: *transactionService}
}

//Route ...
func (controller *TransactionController) Route(group *gin.RouterGroup) {
	// group.GET("/api/transaction/:id", controller.GetOne)
	group.GET("/api/transaction/:warehouse_id", controller.List)
	group.POST("/api/transaction", controller.Create)
	group.DELETE("/api/transaction/:id", controller.Delete)
}

//Create ...
func (controller *TransactionController) Create(context *gin.Context) {
	var request request.CreateTransactionRequest
	context.Bind(&request)

	strToken := context.Request.Header.Get("Authorization")
	token, err := service.VerifyToken(strToken)
	if err != nil {
		log.Println("error --> ", err.Error())
	}

	resp := controller.TransactionService.Create(request, token)

	if resp.ID < 1 {
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

// //GetOne ...
// func (controller *TransactionController) GetOne(context *gin.Context) {
// 	resp := controller.TransactionService.GetById(context.Param("id"))

// 	if resp.Description == "" {
// 		context.JSON(http.StatusOK,
// 			model.WebResponse{
// 				Code:   http.StatusBadRequest,
// 				Status: "Select Failed",
// 				Data:   nil,
// 			})
// 	} else {
// 		context.JSON(http.StatusOK,
// 			model.WebResponse{
// 				Code:   http.StatusOK,
// 				Status: "Select Success",
// 				Data:   resp,
// 			})
// 	}
// }

//List ...
func (controller *TransactionController) List(context *gin.Context) {
	resp := controller.TransactionService.List(context.Param("warehouse_id"))

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
func (controller *TransactionController) Delete(context *gin.Context) {

	controller.TransactionService.DeleteById(context.Param("id"))

	context.JSON(http.StatusOK,
		model.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   nil,
		})
}
