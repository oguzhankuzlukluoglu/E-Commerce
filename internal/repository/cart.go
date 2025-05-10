package repository

import (
	"context"

	"github.com/oguzhan/e-commerce/internal/model"
)

type CartRepository interface {
	GetByUserID(ctx context.Context, userID uint) (*model.Cart, error)
	Create(ctx context.Context, cart *model.Cart) error
	Update(ctx context.Context, cart *model.Cart) error
	Delete(ctx context.Context, userID uint) error
}
