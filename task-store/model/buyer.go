package model

import (
	"github.com/jinzhu/gorm"
)

type Buyer struct {
	gorm.Model
	Name string `json:"name`
}
