package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/oguzhan/e-commerce/internal/model"
	"github.com/oguzhan/e-commerce/pkg/errors"
)

func TestProductService_Create(t *testing.T) {
	mockRepo := new(MockProductRepo)
	service := NewProductService(mockRepo)

	tests := []struct {
		name          string
		product       *model.Product
		mockSetup     func()
		expectedError error
	}{
		{
			name: "successful product creation",
			product: &model.Product{
				Name:        "Test Product",
				Description: "Test Description",
				Price:       99.99,
				Stock:       100,
			},
			mockSetup: func() {
				mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*model.Product")).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "invalid product data",
			product: &model.Product{
				Name:        "", // Empty name
				Description: "Test Description",
				Price:       -10, // Negative price
				Stock:       -5,  // Negative stock
			},
			mockSetup: func() {
				// No mock setup needed
			},
			expectedError: errors.ErrValidation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			err := service.Create(context.Background(), tt.product)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestProductService_GetByID(t *testing.T) {
	mockRepo := new(MockProductRepo)
	service := NewProductService(mockRepo)

	tests := []struct {
		name          string
		productID     uint
		mockSetup     func()
		expectedError error
	}{
		{
			name:      "successful product retrieval",
			productID: 1,
			mockSetup: func() {
				mockRepo.On("GetByID", mock.Anything, uint(1)).Return(&model.Product{
					ID:          1,
					Name:        "Test Product",
					Description: "Test Description",
					Price:       99.99,
					Stock:       100,
				}, nil)
			},
			expectedError: nil,
		},
		{
			name:      "product not found",
			productID: 999,
			mockSetup: func() {
				mockRepo.On("GetByID", mock.Anything, uint(999)).Return(nil, errors.ErrNotFound)
			},
			expectedError: errors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			product, err := service.GetByID(context.Background(), tt.productID)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
				assert.Nil(t, product)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, product)
				assert.Equal(t, tt.productID, product.ID)
			}
		})
	}
}

func TestProductService_List(t *testing.T) {
	mockRepo := new(MockProductRepo)
	service := NewProductService(mockRepo)

	tests := []struct {
		name          string
		offset        int
		limit         int
		mockSetup     func()
		expectedError error
	}{
		{
			name:   "successful product list",
			offset: 0,
			limit:  10,
			mockSetup: func() {
				mockRepo.On("List", mock.Anything, 0, 10).Return([]*model.Product{
					{
						ID:          1,
						Name:        "Product 1",
						Description: "Description 1",
						Price:       99.99,
						Stock:       100,
					},
					{
						ID:          2,
						Name:        "Product 2",
						Description: "Description 2",
						Price:       149.99,
						Stock:       50,
					},
				}, nil)
			},
			expectedError: nil,
		},
		{
			name:   "invalid pagination parameters",
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
			products, err := service.List(context.Background(), tt.offset, tt.limit)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
				assert.Nil(t, products)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, products)
				assert.Len(t, products, 2)
			}
		})
	}
}
