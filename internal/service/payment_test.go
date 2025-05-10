package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/oguzhan/e-commerce/internal/model"
	"github.com/oguzhan/e-commerce/pkg/errors"
)

func TestPaymentService_ProcessPayment(t *testing.T) {
	mockPaymentRepo := new(MockPaymentRepo)
	mockOrderRepo := new(MockOrderRepo)
	service := NewPaymentService(mockPaymentRepo, mockOrderRepo)

	tests := []struct {
		name          string
		orderID       uint
		amount        float64
		paymentMethod string
		mockSetup     func()
		expectedError error
	}{
		{
			name:          "successful payment processing",
			orderID:       1,
			amount:        299.97,
			paymentMethod: "credit_card",
			mockSetup: func() {
				// Mock order retrieval
				mockOrderRepo.On("GetByID", mock.Anything, uint(1)).Return(&model.Order{
					ID:     1,
					UserID: 1,
					Total:  299.97,
					Status: "pending",
				}, nil)

				// Mock payment creation
				mockPaymentRepo.On("Create", mock.Anything, mock.AnythingOfType("*model.Payment")).Return(nil)

				// Mock order update
				mockOrderRepo.On("Update", mock.Anything, mock.AnythingOfType("*model.Order")).Return(nil)
			},
			expectedError: nil,
		},
		{
			name:          "order not found",
			orderID:       999,
			amount:        299.97,
			paymentMethod: "credit_card",
			mockSetup: func() {
				mockOrderRepo.On("GetByID", mock.Anything, uint(999)).Return(nil, errors.ErrNotFound)
			},
			expectedError: errors.ErrNotFound,
		},
		{
			name:          "order already paid",
			orderID:       1,
			amount:        299.97,
			paymentMethod: "credit_card",
			mockSetup: func() {
				mockOrderRepo.On("GetByID", mock.Anything, uint(1)).Return(&model.Order{
					ID:     1,
					UserID: 1,
					Total:  299.97,
					Status: "paid",
				}, nil)
			},
			expectedError: errors.ErrValidation,
		},
		{
			name:          "amount mismatch",
			orderID:       1,
			amount:        200.00, // Different from order total
			paymentMethod: "credit_card",
			mockSetup: func() {
				mockOrderRepo.On("GetByID", mock.Anything, uint(1)).Return(&model.Order{
					ID:     1,
					UserID: 1,
					Total:  299.97,
					Status: "pending",
				}, nil)
			},
			expectedError: errors.ErrValidation,
		},
		{
			name:          "invalid payment method",
			orderID:       1,
			amount:        299.97,
			paymentMethod: "invalid_method",
			mockSetup: func() {
				mockOrderRepo.On("GetByID", mock.Anything, uint(1)).Return(&model.Order{
					ID:     1,
					UserID: 1,
					Total:  299.97,
					Status: "pending",
				}, nil)
			},
			expectedError: errors.ErrValidation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			payment, err := service.ProcessPayment(context.Background(), tt.orderID, tt.amount, tt.paymentMethod)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
				assert.Nil(t, payment)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, payment)
				assert.Equal(t, tt.orderID, payment.OrderID)
				assert.Equal(t, tt.amount, payment.Amount)
				assert.Equal(t, tt.paymentMethod, payment.Method)
				assert.Equal(t, "completed", payment.Status)
			}
		})
	}
}

func TestPaymentService_GetByOrderID(t *testing.T) {
	mockPaymentRepo := new(MockPaymentRepo)
	mockOrderRepo := new(MockOrderRepo)
	service := NewPaymentService(mockPaymentRepo, mockOrderRepo)

	tests := []struct {
		name          string
		orderID       uint
		mockSetup     func()
		expectedError error
	}{
		{
			name:    "successful payment retrieval",
			orderID: 1,
			mockSetup: func() {
				mockPaymentRepo.On("GetByOrderID", mock.Anything, uint(1)).Return(&model.Payment{
					ID:            1,
					OrderID:       1,
					Amount:        299.97,
					Method:        "credit_card",
					Status:        "completed",
					TransactionID: "txn_123",
				}, nil)
			},
			expectedError: nil,
		},
		{
			name:    "payment not found",
			orderID: 999,
			mockSetup: func() {
				mockPaymentRepo.On("GetByOrderID", mock.Anything, uint(999)).Return(nil, errors.ErrNotFound)
			},
			expectedError: errors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			payment, err := service.GetByOrderID(context.Background(), tt.orderID)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
				assert.Nil(t, payment)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, payment)
				assert.Equal(t, tt.orderID, payment.OrderID)
			}
		})
	}
}
