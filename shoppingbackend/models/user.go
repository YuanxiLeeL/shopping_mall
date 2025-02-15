package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null"`
	PhoneNum string `gorm:"type:varchar(255);not null"`
}

type UserWithoutPassword struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Email    string `gorm:"type:varchar(255);not null"`
	PhoneNum string `gorm:"type:varchar(255);not null"`
}
