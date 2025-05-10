package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/oguzhan/e-commerce/internal/model"
	"github.com/oguzhan/e-commerce/pkg/errors"
)

func TestUserService_GetProfile(t *testing.T) {
	mockRepo := new(MockUserRepo)
	service := NewUserService(mockRepo)

	tests := []struct {
		name          string
		userID        uint
		mockSetup     func()
		expectedError error
	}{
		{
			name:   "successful profile retrieval",
			userID: 1,
			mockSetup: func() {
				mockRepo.On("GetByID", mock.Anything, uint(1)).Return(&model.User{
					ID:    1,
					Email: "test@example.com",
				}, nil)
			},
			expectedError: nil,
		},
		{
			name:   "user not found",
			userID: 999,
			mockSetup: func() {
				mockRepo.On("GetByID", mock.Anything, uint(999)).Return(nil, errors.ErrNotFound)
			},
			expectedError: errors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			user, err := service.GetProfile(context.Background(), tt.userID)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.userID, user.ID)
			}
		})
	}
}

func TestUserService_UpdateProfile(t *testing.T) {
	mockRepo := new(MockUserRepo)
	service := NewUserService(mockRepo)

	tests := []struct {
		name          string
		userID        uint
		updateData    map[string]interface{}
		mockSetup     func()
		expectedError error
	}{
		{
			name:   "successful profile update",
			userID: 1,
			updateData: map[string]interface{}{
				"name": "Updated Name",
			},
			mockSetup: func() {
				mockRepo.On("GetByID", mock.Anything, uint(1)).Return(&model.User{
					ID:    1,
					Email: "test@example.com",
				}, nil)
				mockRepo.On("Update", mock.Anything, mock.AnythingOfType("*model.User")).Return(nil)
			},
			expectedError: nil,
		},
		{
			name:   "user not found",
			userID: 999,
			updateData: map[string]interface{}{
				"name": "Updated Name",
			},
			mockSetup: func() {
				mockRepo.On("GetByID", mock.Anything, uint(999)).Return(nil, errors.ErrNotFound)
			},
			expectedError: errors.ErrNotFound,
		},
		{
			name:   "invalid update data",
			userID: 1,
			updateData: map[string]interface{}{
				"invalid_field": "value",
			},
			mockSetup: func() {
				mockRepo.On("GetByID", mock.Anything, uint(1)).Return(&model.User{
					ID:    1,
					Email: "test@example.com",
				}, nil)
			},
			expectedError: errors.ErrValidation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			err := service.UpdateProfile(context.Background(), tt.userID, tt.updateData)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
