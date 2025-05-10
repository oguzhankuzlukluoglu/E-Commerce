package service

import (
	"context"

	"github.com/oguzhan/e-commerce/internal/model"
	"github.com/oguzhan/e-commerce/internal/repository"
	"github.com/oguzhan/e-commerce/pkg/errors"
)

type CartService struct {
	cartRepo    repository.CartRepository
	productRepo repository.ProductRepository
}

func NewCartService(cartRepo repository.CartRepository, productRepo repository.ProductRepository) *CartService {
	return &CartService{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (s *CartService) AddItem(ctx context.Context, userID uint, productID uint, quantity int) error {
	if quantity <= 0 {
		return errors.ErrValidation
	}

	product, err := s.productRepo.GetByID(ctx, productID)
	if err != nil {
		return err
	}

	if product.Stock < quantity {
		return errors.ErrValidation
	}

	cart, err := s.cartRepo.GetByUserID(ctx, userID)
	if err != nil {
		if err == errors.ErrNotFound {
			cart = &model.Cart{
				UserID: userID,
				Items:  []model.CartItem{},
			}
			if err := s.cartRepo.Create(ctx, cart); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	found := false
	for i, item := range cart.Items {
		if item.ProductID == productID {
			cart.Items[i].Quantity += quantity
			found = true
			break
		}
	}

	if !found {
		cart.Items = append(cart.Items, model.CartItem{
			ProductID: productID,
			Quantity:  quantity,
		})
	}

	return s.cartRepo.Update(ctx, cart)
}

func (s *CartService) RemoveItem(ctx context.Context, userID uint, productID uint) error {
	cart, err := s.cartRepo.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}

	found := false
	for i, item := range cart.Items {
		if item.ProductID == productID {
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return errors.ErrNotFound
	}

	return s.cartRepo.Update(ctx, cart)
}

func (s *CartService) GetCart(ctx context.Context, userID uint) (*model.Cart, error) {
	return s.cartRepo.GetByUserID(ctx, userID)
}
