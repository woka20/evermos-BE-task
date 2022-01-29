package model

type Order struct {
	// gorm.Model
	ID       uint `gorm:"column:id;not null;AUTO_INCREMENT"`
	BuyerID  int  `json:"buyer_id"`
	Quantity int  `json:"quantity"`
	StoreID  int  `json:"store_id"`
}
