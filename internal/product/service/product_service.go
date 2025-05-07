package service

import (
	"errors"

	"github.com/oguzhan/e-commerce/internal/product/domain"
)

type productService struct {
	repo domain.ProductRepository
}

func NewProductService(repo domain.ProductRepository) domain.ProductService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(product *domain.Product) error {
	if product.Name == "" {
		return errors.New("product name is required")
	}
	if product.Price <= 0 {
		return errors.New("product price must be greater than zero")
	}
	if product.Stock < 0 {
		return errors.New("product stock cannot be negative")
	}

	return s.repo.Create(product)
}

func (s *productService) GetProduct(id uint) (*domain.Product, error) {
	return s.repo.GetByID(id)
}

func (s *productService) GetAllProducts() ([]*domain.Product, error) {
	return s.repo.GetAll()
}

func (s *productService) UpdateProduct(product *domain.Product) error {
	if product.Name == "" {
		return errors.New("product name is required")
	}
	if product.Price <= 0 {
		return errors.New("product price must be greater than zero")
	}
	if product.Stock < 0 {
		return errors.New("product stock cannot be negative")
	}

	existingProduct, err := s.repo.GetByID(product.ID)
	if err != nil {
		return err
	}

	// Update only allowed fields
	existingProduct.Name = product.Name
	existingProduct.Description = product.Description
	existingProduct.Price = product.Price
	existingProduct.Stock = product.Stock
	existingProduct.Category = product.Category

	return s.repo.Update(existingProduct)
}

func (s *productService) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}

func (s *productService) UpdateStock(id uint, quantity int) error {
	product, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	newStock := product.Stock + quantity
	if newStock < 0 {
		return errors.New("insufficient stock")
	}

	return s.repo.UpdateStock(id, quantity)
}
