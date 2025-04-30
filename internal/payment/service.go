package payment

import (
	"github.com/oguzhan/e-commerce/pkg/models"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	repo  *Repository
	cache *redis.Client
}

func NewService(repo *Repository, cache *redis.Client) *Service {
	return &Service{
		repo:  repo,
		cache: cache,
	}
}

func (s *Service) CreatePayment(payment *models.Payment) error {
	return s.repo.Create(payment)
}

func (s *Service) GetPaymentByID(id uint) (*models.Payment, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetPaymentsByUserID(userID uint) ([]models.Payment, error) {
	return s.repo.GetByUserID(userID)
}

func (s *Service) UpdatePayment(payment *models.Payment) error {
	return s.repo.Update(payment)
}

func (s *Service) ProcessPayment(id uint) error {
	payment, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	payment.Status = models.PaymentStatusCompleted
	return s.repo.Update(payment)
}

func (s *Service) RefundPayment(id uint) error {
	payment, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	payment.Status = models.PaymentStatusRefunded
	return s.repo.Update(payment)
}

func (s *Service) GetPaymentsByOrderID(orderID uint) ([]models.Payment, error) {
	var payments []models.Payment
	if err := s.repo.GetByOrderID(orderID, &payments); err != nil {
		return nil, err
	}
	return payments, nil
}

func (s *Service) ListPayments(page, limit int) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var total int64

	offset := (page - 1) * limit

	if err := s.repo.Count(&total); err != nil {
		return nil, 0, err
	}

	if err := s.repo.GetAll(offset, limit, &payments); err != nil {
		return nil, 0, err
	}

	return payments, total, nil
}
