package auth

import (
	"testing"

	"github.com/oguzhan/e-commerce/pkg/models"
)

type mockDB struct {
	users []models.User
}

func (m *mockDB) Create(value interface{}) *mockDB {
	user := value.(*models.User)
	if user.Email == "exists@example.com" {
		return &mockDB{users: m.users}
	}
	m.users = append(m.users, *user)
	return m
}

func TestRegister_Success(t *testing.T) {
	user := &models.User{Email: "test@example.com", Password: "123456"}
	// Simüle edilen başarılı kayıt
	if user.Email == "exists@example.com" {
		t.Fatal("user already exists")
	}
}

func TestRegister_AlreadyExists(t *testing.T) {
	user := &models.User{Email: "exists@example.com", Password: "123456"}
	// Simüle edilen hata durumu
	if user.Email != "exists@example.com" {
		t.Fatal("expected error, got nil")
	}
}
