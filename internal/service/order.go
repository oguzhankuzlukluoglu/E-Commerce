package service

import (
	"context"

	"github.com/oguzhan/e-commerce/internal/model"
	"github.com/oguzhan/e-commerce/internal/repository"
	"github.com/oguzhan/e-commerce/pkg/errors"
)

type OrderService struct {
	orderRepo   repository.OrderRepository
	productRepo repository.ProductRepository
}

func NewOrderService(orderRepo repository.OrderRepository, productRepo repository.ProductRepository) *OrderService {
	return &OrderService{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (s *OrderService) Create(ctx context.Context, userID uint, items []model.OrderItem) (*model.Order, error) {
	// Validate input
	if len(items) == 0 {
		return nil, errors.ErrValidation
	}

	// Create order
	order := &model.Order{
		UserID: userID,
		Items:  items,
		Status: "pending",
	}

	// Calculate total and validate stock
	var total float64
	for i := range order.Items {
		product, err := s.productRepo.GetByID(ctx, order.Items[i].ProductID)
		if err != nil {
			return nil, errors.ErrNotFound
		}

		// Check stock
		if product.Stock < order.Items[i].Quantity {
			return nil, errors.ErrValidation
		}

		// Update stock
		product.Stock -= order.Items[i].Quantity
		if err := s.productRepo.Update(ctx, product); err != nil {
			return nil, errors.ErrDatabase
		}

		// Set price and calculate total
		order.Items[i].Price = product.Price
		total += product.Price * float64(order.Items[i].Quantity)
	}

	order.Total = total

	// Save order
	if err := s.orderRepo.Create(ctx, order); err != nil {
		return nil, errors.ErrDatabase
	}

	return order, nil
}

func (s *OrderService) GetByID(ctx context.Context, id uint) (*model.Order, error) {
	order, err := s.orderRepo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.ErrNotFound
	}
	return order, nil
}

func (s *OrderService) ListByUserID(ctx context.Context, userID uint, offset, limit int) ([]*model.Order, error) {
	// Validate pagination parameters
	if offset < 0 || limit <= 0 {
		return nil, errors.ErrValidation
	}

	orders, err := s.orderRepo.ListByUserID(ctx, userID, offset, limit)
	if err != nil {
		return nil, errors.ErrDatabase
	}

	return orders, nil
}
