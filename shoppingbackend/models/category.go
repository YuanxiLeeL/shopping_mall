package models

import (
	"gorm.io/gorm"
)

// Category 分类模型
type Category struct {
	gorm.Model
	Name            string `json:"name" binding:"required" gorm:"type:varchar(255);uniqueIndex;not null"`
}
