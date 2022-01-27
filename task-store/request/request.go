package request

type OrderRequest struct {
	BuyerID  int              `json:"user_id" binding:"required"`
	Products []ProductRequest `json:"products" binding:"required"`
}

type ProductRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
}

type ProductPost struct {
	ProductID   int    `json:"product_id" binding:"required"`
	ProductName string `json:"product_name" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
	StoreID     uint   `json:"store_id" binding:"required"`
}

type BuyerRequest struct {
	BuyerID   int    `json:"user_id" binding:"required"`
	BuyerName string `json:"user_name" binding:"required"`
}
