package repository

import (
	"context"

	"github.com/oguzhan/e-commerce/internal/model"
)

type ProductRepository interface {
	Create(ctx context.Context, product *model.Product) error
	GetByID(ctx context.Context, id uint) (*model.Product, error)
	List(ctx context.Context, offset, limit int) ([]*model.Product, error)
	Update(ctx context.Context, product *model.Product) error
	Delete(ctx context.Context, id uint) error
}
