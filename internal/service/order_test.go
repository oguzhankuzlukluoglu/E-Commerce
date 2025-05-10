package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/oguzhan/e-commerce/internal/model"
	"github.com/oguzhan/e-commerce/pkg/errors"
)

func TestOrderService_Create(t *testing.T) {
	mockOrderRepo := new(MockOrderRepo)
	mockProductRepo := new(MockProductRepo)
	service := NewOrderService(mockOrderRepo, mockProductRepo)

	tests := []struct {
		name          string
		userID        uint
		items         []model.OrderItem
		mockSetup     func()
		expectedError error
	}{
		{
			name:   "successful order creation",
			userID: 1,
			items: []model.OrderItem{
				{
					ProductID: 1,
					Quantity:  2,
				},
				{
					ProductID: 2,
					Quantity:  1,
				},
			},
			mockSetup: func() {
				// Mock product availability checks
				mockProductRepo.On("GetByID", mock.Anything, uint(1)).Return(&model.Product{
					ID:    1,
					Stock: 10,
					Price: 99.99,
				}, nil)
				mockProductRepo.On("GetByID", mock.Anything, uint(2)).Return(&model.Product{
					ID:    2,
					Stock: 5,
					Price: 149.99,
				}, nil)

				// Mock order creation
				mockOrderRepo.On("Create", mock.Anything, mock.AnythingOfType("*model.Order")).Return(nil)

				// Mock product stock updates
				mockProductRepo.On("Update", mock.Anything, mock.AnythingOfType("*model.Product")).Return(nil).Times(2)
			},
			expectedError: nil,
		},
		{
			name:   "product not found",
			userID: 1,
			items: []model.OrderItem{
				{
					ProductID: 999,
					Quantity:  1,
				},
			},
			mockSetup: func() {
				mockProductRepo.On("GetByID", mock.Anything, uint(999)).Return(nil, errors.ErrNotFound)
			},
			expectedError: errors.ErrNotFound,
		},
		{
			name:   "insufficient stock",
			userID: 1,
			items: []model.OrderItem{
				{
					ProductID: 1,
					Quantity:  20, // Requesting more than available
				},
			},
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
			name:   "empty order",
			userID: 1,
			items:  []model.OrderItem{},
			mockSetup: func() {
				// No mock setup needed
			},
			expectedError: errors.ErrValidation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			order, err := service.Create(context.Background(), tt.userID, tt.items)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
				assert.Nil(t, order)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, order)
				assert.Equal(t, tt.userID, order.UserID)
				assert.Len(t, order.Items, len(tt.items))
			}
		})
	}
}

func TestOrderService_GetByID(t *testing.T) {
	mockOrderRepo := new(MockOrderRepo)
	mockProductRepo := new(MockProductRepo)
	service := NewOrderService(mockOrderRepo, mockProductRepo)

	tests := []struct {
		name          string
		orderID       uint
		mockSetup     func()
		expectedError error
	}{
		{
			name:    "successful order retrieval",
			orderID: 1,
			mockSetup: func() {
				mockOrderRepo.On("GetByID", mock.Anything, uint(1)).Return(&model.Order{
					ID:     1,
					UserID: 1,
					Items: []model.OrderItem{
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
			name:    "order not found",
			orderID: 999,
			mockSetup: func() {
				mockOrderRepo.On("GetByID", mock.Anything, uint(999)).Return(nil, errors.ErrNotFound)
			},
			expectedError: errors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			order, err := service.GetByID(context.Background(), tt.orderID)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
				assert.Nil(t, order)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, order)
				assert.Equal(t, tt.orderID, order.ID)
			}
		})
	}
}

func TestOrderService_ListByUserID(t *testing.T) {
	mockOrderRepo := new(MockOrderRepo)
	mockProductRepo := new(MockProductRepo)
	service := NewOrderService(mockOrderRepo, mockProductRepo)

	tests := []struct {
		name          string
		userID        uint
		offset        int
		limit         int
		mockSetup     func()
		expectedError error
	}{
		{
			name:   "successful order list",
			userID: 1,
			offset: 0,
			limit:  10,
			mockSetup: func() {
				mockOrderRepo.On("ListByUserID", mock.Anything, uint(1), 0, 10).Return([]*model.Order{
					{
						ID:     1,
						UserID: 1,
						Items: []model.OrderItem{
							{
								ProductID: 1,
								Quantity:  2,
							},
						},
					},
					{
						ID:     2,
						UserID: 1,
						Items: []model.OrderItem{
							{
								ProductID: 2,
								Quantity:  1,
							},
						},
					},
				}, nil)
			},
			expectedError: nil,
		},
		{
			name:   "invalid pagination parameters",
			userID: 1,
			offset: -1,
			limit:  0,
			mockSetup: func() {
				// No mock setup needed
			},
			expectedError: errors.ErrValidation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			orders, err := service.ListByUserID(context.Background(), tt.userID, tt.offset, tt.limit)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
				assert.Nil(t, orders)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, orders)
				assert.Len(t, orders, 2)
			}
		})
	}
}
