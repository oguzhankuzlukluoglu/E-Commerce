package payment

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

func (r *Repository) Create(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

func (r *Repository) GetByID(id uint) (*models.Payment, error) {
	var payment models.Payment
	if err := r.db.Preload("Order").First(&payment, id).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *Repository) GetByUserID(userID uint) ([]models.Payment, error) {
	var payments []models.Payment
	if err := r.db.Preload("Order").Where("user_id = ?", userID).Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

func (r *Repository) Update(payment *models.Payment) error {
	return r.db.Save(payment).Error
}

func (r *Repository) GetByOrderID(orderID uint, payments *[]models.Payment) error {
	return r.db.Preload("Order").
		Where("order_id = ?", orderID).
		Order("created_at DESC").
		Find(payments).Error
}

func (r *Repository) Count(total *int64) error {
	return r.db.Model(&models.Payment{}).Count(total).Error
}

func (r *Repository) GetAll(offset, limit int, payments *[]models.Payment) error {
	return r.db.Preload("Order").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(payments).Error
}
