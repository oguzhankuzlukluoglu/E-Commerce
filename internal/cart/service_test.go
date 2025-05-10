package cart

import (
	"testing"
)

type mockCartDB struct {
	carts []Cart
}

func (m *mockCartDB) FindByUserID(userID uint) *Cart {
	for _, c := range m.carts {
		if c.UserID == userID {
			return &c
		}
	}
	return nil
}

func TestFindByUserID_Found(t *testing.T) {
	db := &mockCartDB{carts: []Cart{{ID: 1, UserID: 10}}}
	cart := db.FindByUserID(10)
	if cart == nil {
		t.Fatal("expected cart, got nil")
	}
}

func TestFindByUserID_NotFound(t *testing.T) {
	db := &mockCartDB{carts: []Cart{{ID: 1, UserID: 10}}}
	cart := db.FindByUserID(20)
	if cart != nil {
		t.Fatal("expected nil, got cart")
	}
}
