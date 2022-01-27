package databases

import (
	// "evermos-be-task/task-store/databases"
	"errors"
	"evermos-be-task/task-store/model"
	"evermos-be-task/task-store/request"
	response "evermos-be-task/task-store/responses"
)

type ProductRepoInterface interface {
	GetProductList() (prods []response.ProductResponse, err error)
	AddProduct(prods model.Product) (err error)
	UpdateProduct(product model.Product) (err error)
	GetProduct(req request.ProductRequest) (prods model.Product, err error)
}

type ProductRepo struct {
}

func NewProductRepo() ProductRepoInterface {
	return &ProductRepo{}

}

func (p *ProductRepo) AddProduct(prods model.Product) (err error) {
	products := &model.Product{ProductName: prods.ProductName, Stock: prods.Stock, StoreID: int(prods.StoreID)}
	database.DB.Create(&products)

	return
}

func (p *ProductRepo) GetProduct(req request.ProductRequest) (prods model.Product, err error) {
	var product model.Product
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.RollbackUnlessCommitted()
		}
	}()
	tx.Where("id = ? and stock >= ?", req.ProductID, req.Quantity).Find(&product)

	if req.ProductID != 0 && req.ProductID == product.ID {
		return product, nil
	} else {
		return model.Product{}, errors.New("Product Not Found")
	}

}

func (p *ProductRepo) GetProductList() (prods []response.ProductResponse, err error) {
	var products []model.Product
	var stores []model.Store
	var storeId []int
	var productResponses []response.ProductResponse
	mapStores := make(map[int]model.Store)

	database.DB.Find(&products)
	for _, val := range products {
		storeId = append(storeId, val.StoreID)
	}

	database.DB.Where("id IN (?)", storeId).Find(&stores)

	for _, val := range stores {
		mapStores[val.StoreID] = val
	}

	for _, product := range products {
		storeResponse := response.StoreResponse{StoreName: mapStores[product.StoreID].StoreName}

		productResponse := response.ProductResponse{
			ID:           int(product.ID),
			ProductName:  product.ProductName,
			CurrentStock: product.Stock,
			Store:        storeResponse,
		}
		productResponses = append(productResponses, productResponse)
	}

	return productResponses, nil

}

func (p *ProductRepo) UpdateProduct(product model.Product) (err error) {
	tx := database.DB.Begin()
	// for _, val := range productList {
	if product.Stock >= 0 {
		if err := tx.Model(&product).Where("id = ?", product.ID).Updates(&product).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	//
	// }
	return
}
