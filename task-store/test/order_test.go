package test

import (
	"bytes"
	"evermos-be-task/task-store/handler"
	response "evermos-be-task/task-store/responses"
	mock_service "evermos-be-task/task-store/test/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
	"github.com/golang/mock/gomock"
	// "github.com/golang/mock/gomock"
	// "github.com/golang/mock/gomock"
	// "github.com/stretchr/testify/assert"
)

const (
	url = "http://localhost:8080/order"
)

func TestOrderConcurrency(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orderSvc := mock_service.NewMockOrder(ctrl)
	orderCtrl := handler.OrderHandler{OrderLogic: orderSvc}

	t.Run("Order RUN", func(t *testing.T) {
		orderSvc.EXPECT().CreateOrderAndOrderDetail(gomock.Eq("{\"user_id\":1,\"products\":[{\"product_id\":1, \"quantity\":2}]}")).Return(response.OrderResponse{}, nil)
		req := httptest.NewRequest("POST", "/order", bytes.NewBufferString(
			"{\"user_id\":1,\"products\":[{\"product_id\":1, \"quantity\":2}]}"))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rec)
		ctx.Request = req

		orderCtrl.GenerateOrder(ctx)
		assert.Equal(t, rec.Code, http.StatusOK)
	})

}
