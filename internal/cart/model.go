package cart

import (
	"time"
)

type Cart struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	UserID    uint       `json:"user_id" gorm:"uniqueIndex"`
	Items     []CartItem `json:"items" gorm:"foreignKey:CartID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type CartItem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CartID    uint      `json:"cart_id"`
	ProductID uint      `json:"product_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
