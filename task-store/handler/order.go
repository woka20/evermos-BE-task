package handler

import (
	"encoding/json"
	"evermos-be-task/task-store/logic"
	"evermos-be-task/task-store/request"
	response "evermos-be-task/task-store/responses"
	"io/ioutil"
	"log"

	// "github.com/kataras/iris/v12"
	"github.com/gin-gonic/gin"
)

type OrderHandlerInterface interface {
	GenerateOrder(ctx *gin.Context)
}

type OrderHandler struct {
	OrderLogic logic.OrderLogicInterface
}

func NewOrderHandler() OrderHandlerInterface {
	return &OrderHandler{
		OrderLogic: logic.NewOrderLogic(),
	}

}

func (o *OrderHandler) GenerateOrder(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(err)
		ctx.Status(500)
		ctx.JSON(500, response.BadResp{
			Status:  500,
			Message: "Error reading request body information",
		})
		// ctx.JSON(response.BadResp{
		// 	Status:  500,
		// 	Message: "Error reading request body information",
		// })
		return
	}

	var order request.OrderRequest

	if err := json.Unmarshal(body, &order); err != nil {
		log.Println(err)
		ctx.Status(500)
		ctx.JSON(500, response.BadResp{
			Status:  500,
			Message: "Error unmarshal",
		})

		return
	}

	OrderResp, err := o.OrderLogic.CreateOrderAndOrderDetail(order)
	if err != nil {
		ctx.JSON(500, response.BadResp{
			Status:  500,
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(200, response.SuccessResp{
		Status: 200,
		Data:   OrderResp,
	})

}
