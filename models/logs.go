package models

import (
	"gorm.io/gorm"
	"time"
)

type Request struct {
	gorm.Model
	BaseCurrency   string    `json:"base_currency"`
	Amount         float64   `json:"amount"`
	TargetCurrency string    `json:"target_currency"`
	CreatedAt      time.Time `json:"created_at"`
}

type Response struct {
	gorm.Model
	RequestID       uint      `json:"request_id"`
	Request         Request   `json:"request" gorm:"foreignKey:RequestID" json:"request"`
	ConvertedAmount float64   `json:"converted_amount"`
	CreatedAt       time.Time `json:"created_at"`
}
