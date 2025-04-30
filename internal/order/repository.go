package order

import (
	"github.com/oguzhan/e-commerce/pkg/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *Repository) GetByID(id uint) (*models.Order, error) {
	var order models.Order
	if err := r.db.Preload("OrderItems").First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *Repository) GetByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	if err := r.db.Preload("OrderItems").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *Repository) Update(order *models.Order) error {
	return r.db.Save(order).Error
}

func (r *Repository) UpdateStatus(id uint, status models.OrderStatus) error {
	return r.db.Model(&models.Order{}).Where("id = ?", id).Update("status", status).Error
}

func (r *Repository) Cancel(order *models.Order) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Order{}).Where("id = ?", order.ID).
			Update("status", models.OrderStatusCancelled).Error; err != nil {
			return err
		}

		for _, item := range order.OrderItems {
			if err := tx.Model(&models.Product{}).Where("id = ?", item.ProductID).
				Update("stock", gorm.Expr("stock + ?", item.Quantity)).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *Repository) List(page, limit int) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64

	offset := (page - 1) * limit

	if err := r.db.Model(&models.Order{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Preload("OrderItems").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}
