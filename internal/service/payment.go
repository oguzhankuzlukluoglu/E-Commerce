package service

import (
	"context"

	"github.com/oguzhan/e-commerce/internal/model"
	"github.com/oguzhan/e-commerce/internal/repository"
	"github.com/oguzhan/e-commerce/pkg/errors"
)

type PaymentService struct {
	paymentRepo repository.PaymentRepository
	orderRepo   repository.OrderRepository
}

func NewPaymentService(paymentRepo repository.PaymentRepository, orderRepo repository.OrderRepository) *PaymentService {
	return &PaymentService{
		paymentRepo: paymentRepo,
		orderRepo:   orderRepo,
	}
}

func (s *PaymentService) ProcessPayment(ctx context.Context, orderID uint, amount float64, paymentMethod string) (*model.Payment, error) {
	order, err := s.orderRepo.GetByID(ctx, orderID)
	if err != nil {
		return nil, err
	}

	if order.Status == "paid" {
		return nil, errors.ErrValidation
	}

	if order.Total != amount {
		return nil, errors.ErrValidation
	}

	if paymentMethod != "credit_card" && paymentMethod != "paypal" {
		return nil, errors.ErrValidation
	}

	payment := &model.Payment{
		OrderID: orderID,
		Amount:  amount,
		Method:  paymentMethod,
		Status:  "completed",
	}

	if err := s.paymentRepo.Create(ctx, payment); err != nil {
		return nil, err
	}

	order.Status = "paid"
	if err := s.orderRepo.Update(ctx, order); err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *PaymentService) GetByOrderID(ctx context.Context, orderID uint) (*model.Payment, error) {
	return s.paymentRepo.GetByOrderID(ctx, orderID)
}
