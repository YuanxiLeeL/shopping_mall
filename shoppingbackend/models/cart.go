package models

import "gorm.io/gorm"

type Cart struct {
	Username string
	gorm.Model
}