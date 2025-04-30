package product

import (
	"errors"

	"github.com/oguzhan/e-commerce/pkg/models"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateProduct(product *models.Product) error {
	return s.db.Create(product).Error
}

func (s *Service) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := s.db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func (s *Service) GetProductBySKU(sku string) (*models.Product, error) {
	var product models.Product
	if err := s.db.Where("sku = ?", sku).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func (s *Service) UpdateProduct(id uint, product *models.Product) error {
	existingProduct, err := s.GetProductByID(id)
	if err != nil {
		return err
	}

	// Update only non-zero fields
	if product.Name != "" {
		existingProduct.Name = product.Name
	}
	if product.Description != "" {
		existingProduct.Description = product.Description
	}
	if product.Price != 0 {
		existingProduct.Price = product.Price
	}
	if product.Stock != 0 {
		existingProduct.Stock = product.Stock
	}
	if product.Category != "" {
		existingProduct.Category = product.Category
	}
	if product.ImageURL != "" {
		existingProduct.ImageURL = product.ImageURL
	}
	if product.SKU != "" {
		existingProduct.SKU = product.SKU
	}
	existingProduct.IsActive = product.IsActive

	return s.db.Save(existingProduct).Error
}

func (s *Service) DeleteProduct(id uint) error {
	return s.db.Delete(&models.Product{}, id).Error
}

func (s *Service) ListProducts(page, limit int) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	offset := (page - 1) * limit

	if err := s.db.Model(&models.Product{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := s.db.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (s *Service) UpdateStock(id uint, quantity int) error {
	return s.db.Model(&models.Product{}).Where("id = ?", id).
		Update("stock", gorm.Expr("stock + ?", quantity)).Error
}

func (s *Service) SearchProducts(query string, category string, minPrice, maxPrice float64) ([]models.Product, error) {
	var products []models.Product
	db := s.db.Model(&models.Product{})

	if query != "" {
		db = db.Where("name ILIKE ? OR description ILIKE ?", "%"+query+"%", "%"+query+"%")
	}
	if category != "" {
		db = db.Where("category = ?", category)
	}
	if minPrice > 0 {
		db = db.Where("price >= ?", minPrice)
	}
	if maxPrice > 0 {
		db = db.Where("price <= ?", maxPrice)
	}

	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
