package models

import "gorm.io/gorm"

type Good struct {
	gorm.Model
	Name        string `json:"name" binding:"required" gorm:"type:varchar(100);not_null"`
	Description string `json:"description" binding:"required" gorm:"type:text;not_null"`
	Price       int64  `json:"price" binding:"required" gorm:"type:bigint;not_null"`
	Category  string   `json:"category" gorm:"not null"`
}

type Goods struct {
	Goods []Good `json:"goods"`
	Total int64
}
