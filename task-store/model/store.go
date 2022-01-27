package model

import (
	"github.com/jinzhu/gorm"
)

type Store struct {
	gorm.Model
	StoreID   int    `json:"store_id"`
	StoreName string `json:"store_name"`
}
