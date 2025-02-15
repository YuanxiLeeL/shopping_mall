package models

import "time"

type ExchangeRate struct {
	ID           uint      `gorm:"primaryKey" json:"_id"`
	FromCurrency string    `gorm:"type:varchar(3)" json:"fromCurrency" binding:"required"`
	ToCurrency   string    `gorm:"type:varchar(3)" json:"toCurrency" binding:"required"`
	Rate         float64   `json:"rate" binding:"required"`
	Date         time.Time `json:"date"`
}
