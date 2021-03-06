package model

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	ProductName string `json:"product_name"`
	Stock       int    `json:"stock" validate: min=0`
	StoreID     int    `json:"store_id"`
}
