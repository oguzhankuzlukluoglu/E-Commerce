package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/oguzhan/e-commerce/internal/model"
	"github.com/oguzhan/e-commerce/pkg/errors"
)

func TestCartService_AddItem(t *testing.T) {
	mockCartRepo := new(MockCartRepo)
	mockProductRepo := new(MockProductRepo)
	service := NewCartService(mockCartRepo, mockProductRepo)

	tests := []struct {
		name          string
		userID        uint
		productID     uint
		quantity      int
		mockSetup     func()
		expectedError error
	}{
		{
			name:      "successful item addition",
			userID:    1,
			productID: 1,
			quantity:  2,
			mockSetup: func() {
				// Mock product retrieval
				mockProductRepo.On("GetByID", mock.Anything, uint(1)).Return(&model.Product{
					ID:    1,
					Stock: 10,
					Price: 99.99,
				}, nil)

				// Mock cart retrieval
				mockCartRepo.On("GetByUserID", mock.Anything, uint(1)).Return(&model.Cart{
					UserID: 1,
					Items:  []model.CartItem{},
				}, nil)

				// Mock cart update
				mockCartRepo.On("Update", mock.Anything, mock.AnythingOfType("*model.Cart")).Return(nil)
			},
			expectedError: nil,
		},
		{
			name:      "product not found",
			userID:    1,
			productID: 999,
			quantity:  1,
			mockSetup: func() {
				mockProductRepo.On("GetByID", mock.Anything, uint(999)).Return(nil, errors.ErrNotFound)
			},
			expectedError: errors.ErrNotFound,
		},
		{
			name:      "insufficient stock",
			userID:    1,
			productID: 1,
			quantity:  20, // Requesting more than available
			mockSetup: func() {
				mockProductRepo.On("GetByID", mock.Anything, uint(1)).Return(&model.Product{
					ID:    1,
					Stock: 10,
					Price: 99.99,
				}, nil)
			},
			expectedError: errors.ErrValidation,
		},
		{
			name:      "invalid quantity",
			userID:    1,
			productID: 1,
			quantity:  0, // Invalid quantity
			mockSetup: func() {
				mockProductRepo.On("GetByID", mock.Anything, uint(1)).Return(&model.Product{
					ID:    1,
					Stock: 10,
					Price: 99.99,
				}, nil)
			},
			expectedError: errors.ErrValidation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			err := service.AddItem(context.Background(), tt.userID, tt.productID, tt.quantity)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCartService_RemoveItem(t *testing.T) {
	mockCartRepo := new(MockCartRepo)
	mockProductRepo := new(MockProductRepo)
	service := NewCartService(mockCartRepo, mockProductRepo)

	tests := []struct {
		name          string
		userID        uint
		productID     uint
		mockSetup     func()
		expectedError error
	}{
		{
			name:      "successful item removal",
			userID:    1,
			productID: 1,
			mockSetup: func() {
				mockCartRepo.On("GetByUserID", mock.Anything, uint(1)).Return(&model.Cart{
					UserID: 1,
					Items: []model.CartItem{
						{
							ProductID: 1,
							Quantity:  2,
						},
					},
				}, nil)
				mockCartRepo.On("Update", mock.Anything, mock.AnythingOfType("*model.Cart")).Return(nil)
			},
			expectedError: nil,
		},
		{
			name:      "item not in cart",
			userID:    1,
			productID: 999,
			mockSetup: func() {
				mockCartRepo.On("GetByUserID", mock.Anything, uint(1)).Return(&model.Cart{
					UserID: 1,
					Items: []model.CartItem{
						{
							ProductID: 1,
							Quantity:  2,
						},
					},
				}, nil)
			},
			expectedError: errors.ErrNotFound,
		},
		{
			name:      "cart not found",
			userID:    999,
			productID: 1,
			mockSetup: func() {
				mockCartRepo.On("GetByUserID", mock.Anything, uint(999)).Return(nil, errors.ErrNotFound)
			},
			expectedError: errors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			err := service.RemoveItem(context.Background(), tt.userID, tt.productID)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCartService_GetCart(t *testing.T) {
	mockCartRepo := new(MockCartRepo)
	mockProductRepo := new(MockProductRepo)
	service := NewCartService(mockCartRepo, mockProductRepo)

	tests := []struct {
		name          string
		userID        uint
		mockSetup     func()
		expectedError error
	}{
		{
			name:   "successful cart retrieval",
			userID: 1,
			mockSetup: func() {
				mockCartRepo.On("GetByUserID", mock.Anything, uint(1)).Return(&model.Cart{
					UserID: 1,
					Items: []model.CartItem{
						{
							ProductID: 1,
							Quantity:  2,
						},
					},
				}, nil)
			},
			expectedError: nil,
		},
		{
			name:   "cart not found",
			userID: 999,
			mockSetup: func() {
				mockCartRepo.On("GetByUserID", mock.Anything, uint(999)).Return(nil, errors.ErrNotFound)
			},
			expectedError: errors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			cart, err := service.GetCart(context.Background(), tt.userID)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
				assert.Nil(t, cart)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, cart)
				assert.Equal(t, tt.userID, cart.UserID)
			}
		})
	}
}
