package models

import (
	"time"

	"gorm.io/gorm"
)

type PaymentStatus string

const (
	PaymentStatusPending    PaymentStatus = "pending"
	PaymentStatusProcessing PaymentStatus = "processing"
	PaymentStatusCompleted  PaymentStatus = "completed"
	PaymentStatusFailed     PaymentStatus = "failed"
	PaymentStatusRefunded   PaymentStatus = "refunded"
)

type Payment struct {
	gorm.Model
	OrderID       uint          `gorm:"not null" json:"order_id"`
	UserID        uint          `gorm:"not null" json:"user_id"`
	Amount        float64       `gorm:"not null" json:"amount"`
	Status        PaymentStatus `gorm:"type:varchar(20);default:'pending'" json:"status"`
	PaymentMethod string        `gorm:"not null" json:"payment_method"`
	TransactionID string        `gorm:"uniqueIndex" json:"transaction_id"`
	PaymentDate   time.Time     `json:"payment_date"`
	Order         Order         `json:"order"`
}

type PaymentResponse struct {
	ID            uint          `json:"id"`
	OrderID       uint          `json:"order_id"`
	UserID        uint          `json:"user_id"`
	Amount        float64       `json:"amount"`
	Status        PaymentStatus `json:"status"`
	PaymentMethod string        `json:"payment_method"`
	TransactionID string        `json:"transaction_id"`
	PaymentDate   time.Time     `json:"payment_date"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}
