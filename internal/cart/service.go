package cart

import (
	"errors"

	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) GetCartByUserID(userID uint) (*Cart, error) {
	var cart Cart
	err := s.db.Preload("Items").Where("user_id = ?", userID).First(&cart).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		cart = Cart{UserID: userID}
		s.db.Create(&cart)
		return &cart, nil
	}
	return &cart, err
}

func (s *Service) AddItem(userID, productID uint, quantity int) error {
	cart, err := s.GetCartByUserID(userID)
	if err != nil {
		return err
	}
	var item CartItem
	err = s.db.Where("cart_id = ? AND product_id = ?", cart.ID, productID).First(&item).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		item = CartItem{CartID: cart.ID, ProductID: productID, Quantity: quantity}
		return s.db.Create(&item).Error
	}
	item.Quantity += quantity
	return s.db.Save(&item).Error
}

func (s *Service) UpdateItem(userID, itemID uint, quantity int) error {
	cart, err := s.GetCartByUserID(userID)
	if err != nil {
		return err
	}
	var item CartItem
	err = s.db.Where("id = ? AND cart_id = ?", itemID, cart.ID).First(&item).Error
	if err != nil {
		return err
	}
	item.Quantity = quantity
	return s.db.Save(&item).Error
}

func (s *Service) RemoveItem(userID, itemID uint) error {
	cart, err := s.GetCartByUserID(userID)
	if err != nil {
		return err
	}
	return s.db.Where("id = ? AND cart_id = ?", itemID, cart.ID).Delete(&CartItem{}).Error
}

func (s *Service) ClearCart(userID uint) error {
	cart, err := s.GetCartByUserID(userID)
	if err != nil {
		return err
	}
	return s.db.Where("cart_id = ?", cart.ID).Delete(&CartItem{}).Error
}
