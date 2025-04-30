package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

type Order struct {
	gorm.Model
	UserID          uint        `gorm:"not null" json:"user_id"`
	Status          OrderStatus `gorm:"type:varchar(20);default:'pending'" json:"status"`
	TotalAmount     float64     `gorm:"not null" json:"total_amount"`
	ShippingAddress string      `gorm:"not null" json:"shipping_address"`
	BillingAddress  string      `gorm:"not null" json:"billing_address"`
	PaymentMethod   string      `gorm:"not null" json:"payment_method"`
	OrderItems      []OrderItem `json:"order_items"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `gorm:"not null" json:"order_id"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	Price     float64 `gorm:"not null" json:"price"`
	Product   Product `json:"product"`
}

type OrderResponse struct {
	ID              uint        `json:"id"`
	UserID          uint        `json:"user_id"`
	Status          OrderStatus `json:"status"`
	TotalAmount     float64     `json:"total_amount"`
	ShippingAddress string      `json:"shipping_address"`
	BillingAddress  string      `json:"billing_address"`
	PaymentMethod   string      `json:"payment_method"`
	OrderItems      []OrderItem `json:"order_items"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}
