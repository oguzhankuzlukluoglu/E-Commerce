package payment

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

func (s *Service) CreatePayment(payment *models.Payment) error {
	return s.db.Create(payment).Error
}

func (s *Service) ProcessPayment(id, userID uint) error {
	var payment models.Payment
	if err := s.db.First(&payment, id).Error; err != nil {
		return err
	}

	if payment.UserID != userID {
		return errors.New("unauthorized")
	}

	return s.db.Model(&payment).Update("status", "completed").Error
}

func (s *Service) GetPaymentByID(id uint) (*models.Payment, error) {
	var payment models.Payment
	if err := s.db.First(&payment, id).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (s *Service) GetPaymentsByUserID(userID uint) ([]models.Payment, error) {
	var payments []models.Payment
	if err := s.db.Where("user_id = ?", userID).Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

func (s *Service) GetPaymentsByOrderID(orderID uint) ([]models.Payment, error) {
	var payments []models.Payment
	if err := s.db.Where("order_id = ?", orderID).Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

func (s *Service) RefundPayment(id, userID uint) error {
	var payment models.Payment
	if err := s.db.First(&payment, id).Error; err != nil {
		return err
	}

	if payment.UserID != userID {
		return errors.New("unauthorized")
	}

	return s.db.Model(&payment).Update("status", "refunded").Error
}

func (s *Service) ListPayments(page, limit int) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var total int64
	if err := s.db.Limit(limit).Offset((page - 1) * limit).Find(&payments).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return payments, total, nil
}
