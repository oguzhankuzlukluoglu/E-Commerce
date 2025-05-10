package product

import (
	"testing"

	"github.com/oguzhan/e-commerce/pkg/models"
	"gorm.io/gorm"
)

type mockProductDB struct {
	products []models.Product
}

func (m *mockProductDB) FindByID(id uint) *models.Product {
	for _, p := range m.products {
		if p.ID == id {
			return &p
		}
	}
	return nil
}

func TestFindByID_Found(t *testing.T) {
	db := &mockProductDB{products: []models.Product{{Model: gorm.Model{ID: 1}}}}
	product := db.FindByID(1)
	if product == nil {
		t.Fatal("expected product, got nil")
	}
}

func TestFindByID_NotFound(t *testing.T) {
	db := &mockProductDB{products: []models.Product{{Model: gorm.Model{ID: 1}}}}
	product := db.FindByID(2)
	if product != nil {
		t.Fatal("expected nil, got product")
	}
}
