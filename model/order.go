package model

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	OrderID int `json:"order_id`
	BuyerID int `json:"buyer_id`
	Qty     int `json:"qty" validate: min=0`
	StoreID int `json:"store_id"`
}
