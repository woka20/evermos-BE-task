package databases

import (
	// "evermos-be-task/task-store/databases"
	"evermos-be-task/task-store/model"
)

type OrderRepoInterface interface {
	// GetOrderList() (prods []response.OrderResponse, err error)
	AddOrder(ord model.Order) (in uint, err error)
}

type OrderRepo struct {
}

func NewOrderRepo() OrderRepoInterface {
	return &OrderRepo{}

}

func (p *OrderRepo) AddOrder(ord model.Order) (in uint, err error) {
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.RollbackUnlessCommitted()
		}
	}()

	orders := &model.Order{
		BuyerID:  ord.BuyerID,
		Quantity: ord.Quantity,
		StoreID:  ord.StoreID,
	}

	tx.Create(&orders)
	// 	tx.Rollback()
	// 	return err
	// }

	tx.Commit()
	return orders.ID, nil
}

// func (p *OrderRepo) GetOrderList() (prods []response.OrderResponse, err error) {
// 	var products []model.Product
// 	var stores []model.Store
// 	var storeId []int
// 	var productResponses []response.ProductResponse
// 	mapStores := make(map[int]model.Store)

// 	database.DB.Find(&products)
// 	for _, val := range products {
// 		storeId = append(storeId, val.StoreID)
// 	}

// 	database.DB.Where("id IN (?)", storeId).Find(&stores)

// 	for _, val := range stores {
// 		mapStores[val.StoreID] = val
// 	}

// 	for _, product := range products {
// 		storeResponse := response.StoreResponse{StoreName: mapStores[product.StoreID].StoreName}

// 		productResponse := response.ProductResponse{
// 			ID:           product.ProductID,
// 			ProductName:  product.ProductName,
// 			CurrentStock: product.Stock,
// 			Store:        storeResponse,
// 		}
// 		productResponses = append(productResponses, productResponse)
// 	}

// 	return productResponses, nil

// }
