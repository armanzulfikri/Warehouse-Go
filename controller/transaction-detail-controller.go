package controller

import (
	"net/http"
	"warehouse/model"
	"warehouse/model/request"
	"warehouse/model/response"
	"warehouse/service"

	"github.com/gin-gonic/gin"
)

//TransactionDetailController ...
type TransactionDetailController struct {
	TransactionDetailService service.TransactionDetailService
}

//NewTransactionDetailController ...
func NewTransactionDetailController(transactionDetailService *service.TransactionDetailService) TransactionDetailController {
	return TransactionDetailController{TransactionDetailService: *transactionDetailService}
}

//Route ...
func (controller *TransactionDetailController) Route(group *gin.RouterGroup) {
	// group.GET("/api/transaction/:id", controller.GetOne)
	// group.GET("/api/detail/transaction", controller.List)
	group.GET("/api/detail/transaction", controller.List)
	group.POST("/api/detail/transaction", controller.Create)
	group.DELETE("/api/detail/:id", controller.Delete)
}

//Create ...
func (controller *TransactionDetailController) Create(context *gin.Context) {
	var request request.CreateDetailRequest
	context.Bind(&request)

	resp := controller.TransactionDetailService.Create(request)

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

//List ...
func (controller *TransactionDetailController) List(context *gin.Context) {
	warehouseID := context.Query("warehouse_id")
	transID := context.Query("transaction_id")
	var resp []response.DetailResponse

	if warehouseID != "" && transID == "" {
		resp = controller.TransactionDetailService.ListByWarehouse(warehouseID) //warehouse id
	}
	if transID != "" && warehouseID == "" {
		resp = controller.TransactionDetailService.ListByTrans(transID)
	}
	if transID == "" && warehouseID == "" {
		resp = controller.TransactionDetailService.ListAll()
	}

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

//ListByTrans ...
// func (controller *TransactionDetailController) ListByTrans(context *gin.Context) {
// 	transactionID := context.Query("transaction_id")
// 	resp := controller.TransactionDetailService.ListByTrans(transactionID) // transaction id

// 	if len(resp) > 0 {
// 		context.JSON(http.StatusOK,
// 			model.WebResponse{
// 				Code:   http.StatusOK,
// 				Status: "OK",
// 				Data:   resp,
// 			})
// 	} else {
// 		context.JSON(http.StatusOK,
// 			model.WebResponse{
// 				Code:   http.StatusOK,
// 				Status: "OK",
// 				Data:   nil,
// 			})
// 	}
// }

//Delete ...
func (controller *TransactionDetailController) Delete(context *gin.Context) {

	controller.TransactionDetailService.DeleteById(context.Param("id"))

	context.JSON(http.StatusOK,
		model.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   nil,
		})
}
