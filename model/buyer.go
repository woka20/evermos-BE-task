package model

import (
	"github.com/jinzhu/gorm"
)

type Buyer struct {
	gorm.Model
	BuyerID int    `json:"buyer_id`
	Name    string `json:"name`
}
