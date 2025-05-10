package repository

import (
	"context"

	"github.com/oguzhan/e-commerce/internal/model"
)

type OrderRepository interface {
	Create(ctx context.Context, order *model.Order) error
	GetByID(ctx context.Context, id uint) (*model.Order, error)
	ListByUserID(ctx context.Context, userID uint, offset, limit int) ([]*model.Order, error)
	Update(ctx context.Context, order *model.Order) error
	Delete(ctx context.Context, id uint) error
}
