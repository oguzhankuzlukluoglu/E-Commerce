package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"

	"github.com/oguzhan/e-commerce/internal/model"
	"github.com/oguzhan/e-commerce/pkg/errors"
)

func TestAuthService_Register(t *testing.T) {
	mockRepo := new(MockUserRepo)
	service := NewAuthService(mockRepo, []byte("test-key"))

	tests := []struct {
		name          string
		email         string
		password      string
		mockSetup     func()
		expectedError error
	}{
		{
			name:     "successful registration",
			email:    "test@example.com",
			password: "password123",
			mockSetup: func() {
				mockRepo.On("GetByEmail", mock.Anything, "test@example.com").Return(nil, errors.ErrNotFound)
				mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*model.User")).Return(nil)
			},
			expectedError: nil,
		},
		{
			name:     "duplicate email",
			email:    "existing@example.com",
			password: "password123",
			mockSetup: func() {
				mockRepo.On("GetByEmail", mock.Anything, "existing@example.com").Return(&model.User{}, nil)
			},
			expectedError: errors.ErrDuplicate,
		},
		{
			name:     "invalid email",
			email:    "invalid-email",
			password: "password123",
			mockSetup: func() {
				// No mock setup needed
			},
			expectedError: errors.ErrValidation,
		},
		{
			name:     "short password",
			email:    "test@example.com",
			password: "short",
			mockSetup: func() {
				// No mock setup needed
			},
			expectedError: errors.ErrValidation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			err := service.Register(context.Background(), tt.email, tt.password)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAuthService_Login(t *testing.T) {
	mockRepo := new(MockUserRepo)
	service := NewAuthService(mockRepo, []byte("test-key"))

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("correctpassword"), bcrypt.DefaultCost)
	user := &model.User{
		Email:    "test@example.com",
		Password: string(hashedPassword),
	}

	tests := []struct {
		name          string
		email         string
		password      string
		mockSetup     func()
		expectedError error
	}{
		{
			name:     "successful login",
			email:    "test@example.com",
			password: "correctpassword",
			mockSetup: func() {
				mockRepo.On("GetByEmail", mock.Anything, "test@example.com").Return(user, nil)
			},
			expectedError: nil,
		},
		{
			name:     "user not found",
			email:    "nonexistent@example.com",
			password: "password123",
			mockSetup: func() {
				mockRepo.On("GetByEmail", mock.Anything, "nonexistent@example.com").Return(nil, errors.ErrNotFound)
			},
			expectedError: errors.ErrUnauthorized,
		},
		{
			name:     "wrong password",
			email:    "test@example.com",
			password: "wrongpassword",
			mockSetup: func() {
				mockRepo.On("GetByEmail", mock.Anything, "test@example.com").Return(user, nil)
			},
			expectedError: errors.ErrUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			token, err := service.Login(context.Background(), tt.email, tt.password)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
				assert.Empty(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
			}
		})
	}
}
