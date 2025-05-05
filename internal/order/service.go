package order

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

func (s *Service) CreateOrder(order *models.Order) error {
	return s.db.Create(order).Error
}

func (s *Service) GetOrderByID(id uint) (*models.Order, error) {
	var order models.Order
	if err := s.db.Preload("OrderItems").First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *Service) GetOrdersByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	if err := s.db.Preload("OrderItems").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *Service) ListOrders(userID uint, page, limit int) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64

	query := s.db.Model(&models.Order{}).Where("user_id = ?", userID)
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Preload("OrderItems").
		Limit(limit).
		Offset((page - 1) * limit).
		Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}

func (s *Service) UpdateOrder(id uint, order *models.Order) error {
	var existingOrder models.Order
	if err := s.db.First(&existingOrder, id).Error; err != nil {
		return err
	}

	return s.db.Model(&existingOrder).Updates(order).Error
}

func (s *Service) CancelOrder(id uint, userID uint) error {
	var order models.Order
	if err := s.db.First(&order, id).Error; err != nil {
		return err
	}

	if order.UserID != userID {
		return errors.New("unauthorized")
	}

	return s.db.Model(&order).Update("status", "cancelled").Error
}

func (s *Service) UpdateOrderStatus(id, userID uint, status string) error {
	var order models.Order
	if err := s.db.First(&order, id).Error; err != nil {
		return err
	}

	if order.UserID != userID {
		return errors.New("unauthorized")
	}

	return s.db.Model(&order).Update("status", status).Error
}
