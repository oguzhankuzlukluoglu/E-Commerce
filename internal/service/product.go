package service

import (
	"context"

	"github.com/oguzhan/e-commerce/internal/model"
	"github.com/oguzhan/e-commerce/internal/repository"
	"github.com/oguzhan/e-commerce/pkg/errors"
)

type ProductService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

func (s *ProductService) Create(ctx context.Context, product *model.Product) error {
	// Validate input
	if product.Name == "" || product.Price <= 0 || product.Stock < 0 {
		return errors.ErrValidation
	}

	if err := s.productRepo.Create(ctx, product); err != nil {
		return errors.ErrDatabase
	}

	return nil
}

func (s *ProductService) GetByID(ctx context.Context, id uint) (*model.Product, error) {
	product, err := s.productRepo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.ErrNotFound
	}
	return product, nil
}

func (s *ProductService) List(ctx context.Context, offset, limit int) ([]*model.Product, error) {
	// Validate pagination parameters
	if offset < 0 || limit <= 0 {
		return nil, errors.ErrValidation
	}

	products, err := s.productRepo.List(ctx, offset, limit)
	if err != nil {
		return nil, errors.ErrDatabase
	}

	return products, nil
}

func (s *ProductService) Update(ctx context.Context, product *model.Product) error {
	// Validate input
	if product.Name == "" || product.Price <= 0 || product.Stock < 0 {
		return errors.ErrValidation
	}

	if err := s.productRepo.Update(ctx, product); err != nil {
		return errors.ErrDatabase
	}

	return nil
}

func (s *ProductService) Delete(ctx context.Context, id uint) error {
	if err := s.productRepo.Delete(ctx, id); err != nil {
		return errors.ErrDatabase
	}
	return nil
}
