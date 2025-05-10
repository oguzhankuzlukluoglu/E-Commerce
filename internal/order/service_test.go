package order

import (
	"testing"

	"github.com/oguzhan/e-commerce/pkg/models"
	"gorm.io/gorm"
)

type mockOrderDB struct {
	orders []models.Order
}

func (m *mockOrderDB) FindByID(id uint) *models.Order {
	for _, o := range m.orders {
		if o.ID == id {
			return &o
		}
	}
	return nil
}

func TestFindByID_Found(t *testing.T) {
	db := &mockOrderDB{orders: []models.Order{{Model: gorm.Model{ID: 1}}}}
	order := db.FindByID(1)
	if order == nil {
		t.Fatal("expected order, got nil")
	}
}

func TestFindByID_NotFound(t *testing.T) {
	db := &mockOrderDB{orders: []models.Order{{Model: gorm.Model{ID: 1}}}}
	order := db.FindByID(2)
	if order != nil {
		t.Fatal("expected nil, got order")
	}
}
