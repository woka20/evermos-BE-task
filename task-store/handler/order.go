package handler

import (
	"encoding/json"
	"evermos-be-task/task-store/logic"
	"evermos-be-task/task-store/request"
	response "evermos-be-task/task-store/responses"
	"log"

	"github.com/kataras/iris/v12"
)

type OrderHandlerInterface interface {
	GenerateOrder(ctx iris.Context)
}

type OrderHandler struct {
	OrderLogic logic.OrderLogicInterface
}

func NewOrderHandler() OrderHandlerInterface {
	return &OrderHandler{
		OrderLogic: logic.NewOrderLogic(),
	}

}

func (o *OrderHandler) GenerateOrder(ctx iris.Context) {
	body, err := ctx.GetBody()
	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		ctx.JSON(response.BadResp{
			Status:  500,
			Message: "Error reading request body information",
		})
		return
	}

	var order request.OrderRequest

	if err := json.Unmarshal(body, &order); err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		ctx.JSON(response.BadResp{
			Status:  500,
			Message: "Error Umarshal request body information",
		})
		return
	}

	OrderResp, err := o.OrderLogic.CreateOrderAndOrderDetail(order)
	if err != nil {
		ctx.JSON(response.BadResp{
			Status:  500,
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(response.SuccessResp{
		Status: 200,
		Data:   OrderResp,
	})

}
