package user

import (
	"testing"

	"github.com/oguzhan/e-commerce/pkg/models"
)

type mockUserDB struct {
	users []models.User
}

func (m *mockUserDB) FindByEmail(email string) *models.User {
	for _, u := range m.users {
		if u.Email == email {
			return &u
		}
	}
	return nil
}

func TestFindByEmail_Found(t *testing.T) {
	db := &mockUserDB{users: []models.User{{Email: "test@example.com"}}}
	user := db.FindByEmail("test@example.com")
	if user == nil {
		t.Fatal("expected user, got nil")
	}
}

func TestFindByEmail_NotFound(t *testing.T) {
	db := &mockUserDB{users: []models.User{{Email: "test@example.com"}}}
	user := db.FindByEmail("notfound@example.com")
	if user != nil {
		t.Fatal("expected nil, got user")
	}
}
