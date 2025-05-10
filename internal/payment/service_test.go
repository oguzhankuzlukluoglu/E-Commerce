package payment

import (
	"testing"

	"github.com/oguzhan/e-commerce/pkg/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate models
	err = db.AutoMigrate(&models.Order{}, &models.Payment{})
	if err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	// Create test order
	order := &models.Order{
		UserID:      1,
		TotalAmount: 100.00,
		Status:      models.OrderStatusPending,
	}
	if err := db.Create(order).Error; err != nil {
		t.Fatalf("Failed to create test order: %v", err)
	}

	return db
}

func TestCreatePayment(t *testing.T) {
	db := setupTestDB(t)
	service := NewService(db)

	payment := &models.Payment{
		OrderID:       1,
		UserID:        1,
		Amount:        100.00,
		PaymentMethod: "credit_card",
	}

	err := service.CreatePayment(payment)
	assert.NoError(t, err)
	assert.NotZero(t, payment.ID)
	assert.Equal(t, models.PaymentStatusPending, payment.Status)
	assert.NotEmpty(t, payment.TransactionID)
}

func TestProcessPayment(t *testing.T) {
	db := setupTestDB(t)
	service := NewService(db)

	// Create a payment first
	payment := &models.Payment{
		OrderID:       1,
		UserID:        1,
		Amount:        100.00,
		PaymentMethod: "credit_card",
	}
	err := service.CreatePayment(payment)
	assert.NoError(t, err)

	// Process the payment
	err = service.ProcessPayment(payment.ID, payment.UserID)
	assert.NoError(t, err)

	// Verify payment status
	updatedPayment, err := service.GetPaymentByID(payment.ID)
	assert.NoError(t, err)
	assert.Equal(t, models.PaymentStatusCompleted, updatedPayment.Status)
}

func TestGetPaymentByID(t *testing.T) {
	db := setupTestDB(t)
	service := NewService(db)

	// Create a payment
	payment := &models.Payment{
		OrderID:       1,
		UserID:        1,
		Amount:        100.00,
		PaymentMethod: "credit_card",
	}
	err := service.CreatePayment(payment)
	assert.NoError(t, err)

	// Get the payment
	retrievedPayment, err := service.GetPaymentByID(payment.ID)
	assert.NoError(t, err)
	assert.Equal(t, payment.ID, retrievedPayment.ID)
	assert.Equal(t, payment.Amount, retrievedPayment.Amount)
}

func TestGetPaymentsByUserID(t *testing.T) {
	db := setupTestDB(t)
	service := NewService(db)

	// Create another test order
	order2 := &models.Order{
		UserID:      1,
		TotalAmount: 200.00,
		Status:      models.OrderStatusPending,
	}
	if err := db.Create(order2).Error; err != nil {
		t.Fatalf("Failed to create test order: %v", err)
	}

	// Create multiple payments for the same user
	payments := []*models.Payment{
		{
			OrderID:       1,
			UserID:        1,
			Amount:        100.00,
			PaymentMethod: "credit_card",
		},
		{
			OrderID:       2,
			UserID:        1,
			Amount:        200.00,
			PaymentMethod: "credit_card",
		},
	}

	for _, payment := range payments {
		err := service.CreatePayment(payment)
		assert.NoError(t, err)
	}

	// Get payments by user ID
	retrievedPayments, err := service.GetPaymentsByUserID(1)
	assert.NoError(t, err)
	assert.Len(t, retrievedPayments, 2)
}

func TestRefundPayment(t *testing.T) {
	db := setupTestDB(t)
	service := NewService(db)

	// Create and process a payment
	payment := &models.Payment{
		OrderID:       1,
		UserID:        1,
		Amount:        100.00,
		PaymentMethod: "credit_card",
	}
	err := service.CreatePayment(payment)
	assert.NoError(t, err)

	err = service.ProcessPayment(payment.ID, payment.UserID)
	assert.NoError(t, err)

	// Refund the payment
	err = service.RefundPayment(payment.ID, payment.UserID)
	assert.NoError(t, err)

	// Verify payment status
	updatedPayment, err := service.GetPaymentByID(payment.ID)
	assert.NoError(t, err)
	assert.Equal(t, models.PaymentStatusRefunded, updatedPayment.Status)
}

func TestListPayments(t *testing.T) {
	db := setupTestDB(t)
	service := NewService(db)

	// Create a payment for the first order
	payment := &models.Payment{
		OrderID:       1,
		UserID:        1,
		Amount:        100.00,
		PaymentMethod: "credit_card",
	}
	err := service.CreatePayment(payment)
	assert.NoError(t, err)

	// Create multiple orders and payments
	for i := 2; i <= 15; i++ {
		order := &models.Order{
			UserID:      1,
			TotalAmount: float64(i * 100),
			Status:      models.OrderStatusPending,
		}
		if err := db.Create(order).Error; err != nil {
			t.Fatalf("Failed to create test order: %v", err)
		}

		payment := &models.Payment{
			OrderID:       uint(i),
			UserID:        1,
			Amount:        float64(i * 100),
			PaymentMethod: "credit_card",
		}
		err = service.CreatePayment(payment)
		assert.NoError(t, err)
	}

	// Test pagination
	payments, total, err := service.ListPayments(1, 10)
	assert.NoError(t, err)
	assert.Equal(t, int64(15), total)
	assert.Len(t, payments, 10)

	// Test second page
	payments, total, err = service.ListPayments(2, 10)
	assert.NoError(t, err)
	assert.Equal(t, int64(15), total)
	assert.Len(t, payments, 5)
}

type mockPaymentDB struct {
	payments []models.Payment
}

func (m *mockPaymentDB) FindByID(id uint) *models.Payment {
	for _, p := range m.payments {
		if p.ID == id {
			return &p
		}
	}
	return nil
}

func TestFindByID_Found(t *testing.T) {
	db := &mockPaymentDB{payments: []models.Payment{{Model: gorm.Model{ID: 1}}}}
	payment := db.FindByID(1)
	if payment == nil {
		t.Fatal("expected payment, got nil")
	}
}

func TestFindByID_NotFound(t *testing.T) {
	db := &mockPaymentDB{payments: []models.Payment{{Model: gorm.Model{ID: 1}}}}
	payment := db.FindByID(2)
	if payment != nil {
		t.Fatal("expected nil, got payment")
	}
}
