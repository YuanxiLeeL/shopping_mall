package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	Username   string `json:"username" binding:"required" gorm:"index"`
	GoodName   string `json:"goodname" binding:"required" gorm:"index"`
	Quantity int  `json:"quantity" binding:"required,min=1"`
	Price int64 `json:"price"`
}
