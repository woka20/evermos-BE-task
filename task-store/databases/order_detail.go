package databases

import (
	// "evermos-be-task/task-store/databases"
	"evermos-be-task/task-store/model"
)

type OrderDetailRepoInterface interface {
	// GetOrderList() (prods []response.OrderResponse, err error)
	AddOrderDetail(ods model.OrderDetail) (err error)
}

type OrderDetailRepo struct {
}

func NewOrderDetailRepo() OrderDetailRepoInterface {
	return &OrderDetailRepo{}

}

func (p *OrderDetailRepo) AddOrderDetail(ods model.OrderDetail) (err error) {
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.RollbackUnlessCommitted()
		}
	}()

	orderDetails := &model.OrderDetail{
		OrderID:   ods.OrderID,
		ProductID: ods.ProductID,
		Qty:       ods.Qty,
	}
	if err := tx.Create(&orderDetails).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return
}
