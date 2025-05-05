package payment

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/oguzhan/e-commerce/pkg/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestRouter(t *testing.T) (*gin.Engine, *gorm.DB) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

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

	service := NewService(db)
	handler := NewHandler(service)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Add auth middleware for testing
	router.Use(func(c *gin.Context) {
		c.Set("user_id", uint(1))
		c.Next()
	})

	router.POST("/payments", handler.CreatePayment)
	router.POST("/payments/:id/process", handler.ProcessPayment)
	router.GET("/payments/:id", handler.GetPayment)
	router.GET("/payments", handler.ListPayments)
	router.POST("/payments/:id/refund", handler.RefundPayment)

	return router, db
}

func TestCreatePaymentHandler(t *testing.T) {
	router, _ := setupTestRouter(t)

	reqBody := map[string]interface{}{
		"order_id":       1,
		"amount":         100.00,
		"payment_method": "credit_card",
	}
	jsonData, _ := json.Marshal(reqBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/payments", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.Payment
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.NotEmpty(t, response.TransactionID)
}

func TestProcessPaymentHandler(t *testing.T) {
	router, _ := setupTestRouter(t)

	// First create a payment
	createReqBody := map[string]interface{}{
		"order_id":       1,
		"amount":         100.00,
		"payment_method": "credit_card",
	}
	jsonData, _ := json.Marshal(createReqBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/payments", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var payment models.Payment
	json.Unmarshal(w.Body.Bytes(), &payment)

	// Then process it
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/payments/"+strconv.Itoa(int(payment.ID))+"/process", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetPaymentHandler(t *testing.T) {
	router, _ := setupTestRouter(t)

	// First create a payment
	createReqBody := map[string]interface{}{
		"order_id":       1,
		"amount":         100.00,
		"payment_method": "credit_card",
	}
	jsonData, _ := json.Marshal(createReqBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/payments", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var payment models.Payment
	json.Unmarshal(w.Body.Bytes(), &payment)

	// Then get it
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/payments/"+strconv.Itoa(int(payment.ID)), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.Payment
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, payment.ID, response.ID)
	assert.Equal(t, payment.Amount, response.Amount)
	assert.NotEmpty(t, response.TransactionID)
}

func TestListPaymentsHandler(t *testing.T) {
	router, db := setupTestRouter(t)

	// Create a payment for the first order
	createReqBody := map[string]interface{}{
		"order_id":       1,
		"amount":         100.00,
		"payment_method": "credit_card",
	}
	jsonData, _ := json.Marshal(createReqBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/payments", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	// Create multiple orders and payments
	for i := 2; i <= 15; i++ {
		amount := float64(i * 100)
		order := &models.Order{
			UserID:      1,
			TotalAmount: amount,
			Status:      models.OrderStatusPending,
		}
		err := db.Create(order).Error
		assert.NoError(t, err)

		createReqBody := map[string]interface{}{
			"order_id":       i,
			"amount":         amount,
			"payment_method": "credit_card",
		}
		jsonData, _ := json.Marshal(createReqBody)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/payments", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	}

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/payments?page=1&limit=10", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	payments := response["payments"].([]interface{})
	assert.Len(t, payments, 10)
	assert.Equal(t, float64(15), response["total"].(float64))
}

func TestRefundPaymentHandler(t *testing.T) {
	router, _ := setupTestRouter(t)

	// First create and process a payment
	createReqBody := map[string]interface{}{
		"order_id":       1,
		"amount":         100.00,
		"payment_method": "credit_card",
	}
	jsonData, _ := json.Marshal(createReqBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/payments", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	var payment models.Payment
	json.Unmarshal(w.Body.Bytes(), &payment)

	// Process the payment
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/payments/"+strconv.Itoa(int(payment.ID))+"/process", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Then refund it
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/payments/"+strconv.Itoa(int(payment.ID))+"/refund", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
