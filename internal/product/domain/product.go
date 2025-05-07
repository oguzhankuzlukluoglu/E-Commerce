package domain

import (
	"time"
)

type Product struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductRepository interface {
	Create(product *Product) error
	GetByID(id uint) (*Product, error)
	GetAll() ([]*Product, error)
	Update(product *Product) error
	Delete(id uint) error
	UpdateStock(id uint, quantity int) error
}

type ProductService interface {
	CreateProduct(product *Product) error
	GetProduct(id uint) (*Product, error)
	GetAllProducts() ([]*Product, error)
	UpdateProduct(product *Product) error
	DeleteProduct(id uint) error
	UpdateStock(id uint, quantity int) error
}
