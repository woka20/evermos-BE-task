package logic

import (
	"errors"
	"evermos-be-task/task-store/databases"
	"evermos-be-task/task-store/model"
	"evermos-be-task/task-store/request"
	response "evermos-be-task/task-store/responses"
	"log"
)

type OrderLogicInterface interface {
	CreateOrderAndOrderDetail(req request.OrderRequest) (resp response.OrderResponse, err error)
}

type OrderLogic struct {
	OrderRepo       databases.OrderRepoInterface
	OrderDetailRepo databases.OrderDetailRepoInterface
	ProductRepo     databases.ProductRepoInterface
}

func NewOrderLogic() OrderLogicInterface {
	return &OrderLogic{
		OrderRepo:       databases.NewOrderRepo(),
		ProductRepo:     databases.NewProductRepo(),
		OrderDetailRepo: databases.NewOrderDetailRepo(),
	}

}

func (o *OrderLogic) CreateOrderAndOrderDetail(req request.OrderRequest) (resp response.OrderResponse, err error) {

	storeIds := make(map[int]int)
	var products []model.Product
	for _, val := range req.Products {
		res, err := o.ProductRepo.GetProduct(val)

		if err != nil {
			log.Println(err)
			return resp, err
		}

		products = append(products, res)
		storeIds[int(res.StoreID)] = res.StoreID

	}

	if len(storeIds) > 1 {
		return response.OrderResponse{}, errors.New("Cannot from different store")
	}

	if len(products) == 0 {
		return response.OrderResponse{}, errors.New("Product Not Available")
	} else if len(products) != len(req.Products) {
		return response.OrderResponse{}, errors.New("Some Product Not Available")
	}

	mapForProduct := mappingProductDetail(products)
	var orderDetailList []model.OrderDetail
	orders := model.Order{
		BuyerID:  req.BuyerID,
		Quantity: countAll(req.Products),
		StoreID:  mapForProduct[0].StoreID,
	}

	o.OrderRepo.AddOrder(orders)

	for _, val := range req.Products {
		orderDetails := model.OrderDetail{
			OrderID:   int(orders.ID),
			ProductID: int(val.ProductID),
			Qty:       val.Quantity,
		}
		o.OrderDetailRepo.AddOrderDetail(orderDetails)
		orderDetailList = append(orderDetailList, orderDetails)
		productUpdate := mapForProduct[int(val.ProductID)]
		productUpdate.Stock = productUpdate.Stock - orderDetails.Qty
		mapForProduct[int(val.ProductID)] = productUpdate

		o.ProductRepo.UpdateProduct(mapForProduct[int(val.ProductID)])

	}

	result := makingOrderResponse(orders, orderDetailList, mapForProduct)

	return result, nil

}

func makingOrderResponse(order model.Order, orderDetail []model.OrderDetail, maps map[int]model.Product) (resp response.OrderResponse) {

	var orderDetailRespList []response.OrderDetailResponse

	for _, val := range orderDetail {
		orderDetailResp := response.OrderDetailResponse{
			OrderDetailID: int(val.ID),
			ProductID:     val.ProductID,
			ProductName:   maps[val.ProductID].ProductName,
			Quantity:      val.Qty,
		}
		orderDetailRespList = append(orderDetailRespList, orderDetailResp)
	}

	orderFinal := response.OrderResponse{
		OrderID:      int(order.ID),
		OrderDetails: orderDetailRespList,
	}

	return orderFinal
}

func mappingProductDetail(products []model.Product) (resp map[int]model.Product) {
	newMap := make(map[int]model.Product)
	for _, val := range products {
		newMap[int(val.ID)] = val

	}
	return newMap

}

func countAll(ods []request.ProductRequest) int {
	var count int
	for _, val := range ods {

		count = count + val.Quantity

	}
	return count
}
