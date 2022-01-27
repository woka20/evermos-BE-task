package logic

import (
	"errors"
	"evermos-be-task/task-store/databases"
	response "evermos-be-task/task-store/responses"
	"log"
)

type ProductLogicInterface interface {
	GetList() (prods []response.ProductResponse, err error)
	// AddOrder(productMap map[int]model.Product, ods request.OrderRequest) (err error)
}

type ProductLogic struct {
	ProductRepo databases.ProductRepoInterface
}

func NewProductLogic() ProductLogicInterface {
	return &ProductLogic{
		ProductRepo: databases.NewProductRepo(),
	}

}

func (p *ProductLogic) GetList() (prods []response.ProductResponse, err error) {
	list, err := p.ProductRepo.GetProductList()

	if err != nil {
		log.Println(err)
		err := errors.New("Failed to fetch List Product")
		return prods, err
	}

	return list, nil

}
