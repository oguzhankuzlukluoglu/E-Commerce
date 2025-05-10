package model

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	OrderID       uint           `gorm:"not null;uniqueIndex" json:"order_id"`
	Amount        float64        `gorm:"not null" json:"amount"`
	Method        string         `gorm:"not null" json:"method"`
	Status        string         `gorm:"not null;default:pending" json:"status"`
	TransactionID string         `gorm:"uniqueIndex" json:"transaction_id"`
}
