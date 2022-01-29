package mock

import (
	"evermos-be-task/task-store/request"
	response "evermos-be-task/task-store/responses"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

type MockOrder struct {
	ctrl     *gomock.Controller
	recorder *MockOrderMockRecorder
}

type MockOrderMockRecorder struct {
	mock *MockOrder
}

func NewMockOrder(ctrl *gomock.Controller) *MockOrder {
	mock := &MockOrder{ctrl: ctrl}
	mock.recorder = &MockOrderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrder) EXPECT() *MockOrderMockRecorder {
	return m.recorder
}

func (m *MockOrder) CreateOrderAndOrderDetail(req request.OrderRequest) (resp response.OrderResponse, err error) {
	m.ctrl.T.Helper()

	ret := m.ctrl.Call(m, "CreateOrderAndOrderDetail", req)
	ret0, _ := ret[0].(response.OrderResponse)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

func (mr *MockOrderMockRecorder) CreateOrderAndOrderDetail(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrderAndOrderDetail", reflect.TypeOf((*MockOrder)(nil).CreateOrderAndOrderDetail), req)
}
