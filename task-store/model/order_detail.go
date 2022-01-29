package model

type OrderDetail struct {
	// gorm.Model
	ID        uint `gorm:"column:id;not null;AUTO_INCREMENT"`
	OrderID   int  `json:"order_id`
	ProductID int  `json:"product_id`
	Qty       int  `json:"qty" validate: min=0`
}
