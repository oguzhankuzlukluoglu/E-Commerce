package order

import (
	"errors"

	"github.com/oguzhan/e-commerce/pkg/models"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateOrder(order *models.Order) error {
	return s.repo.Create(order)
}

func (s *Service) GetOrderByID(id uint) (*models.Order, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetOrdersByUserID(userID uint) ([]models.Order, error) {
	return s.repo.GetByUserID(userID)
}

func (s *Service) UpdateOrder(order *models.Order) error {
	return s.repo.Update(order)
}

func (s *Service) UpdateOrderStatus(id uint, status models.OrderStatus) error {
	return s.repo.UpdateStatus(id, status)
}

func (s *Service) CancelOrder(id uint) error {
	order, err := s.GetOrderByID(id)
	if err != nil {
		return err
	}

	if order.Status == models.OrderStatusCancelled {
		return errors.New("order is already cancelled")
	}

	return s.repo.Cancel(order)
}

func (s *Service) ListOrders(page, limit int) ([]models.Order, int64, error) {
	return s.repo.List(page, limit)
}
