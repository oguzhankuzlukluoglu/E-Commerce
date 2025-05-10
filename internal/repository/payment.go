package repository

import (
	"context"

	"github.com/oguzhan/e-commerce/internal/model"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *model.Payment) error
	GetByID(ctx context.Context, id uint) (*model.Payment, error)
	GetByOrderID(ctx context.Context, orderID uint) (*model.Payment, error)
	Update(ctx context.Context, payment *model.Payment) error
}
