package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/oguzhan/e-commerce/internal/model"
	"github.com/oguzhan/e-commerce/internal/repository"
	"github.com/oguzhan/e-commerce/pkg/errors"
)

type AuthService struct {
	userRepo repository.UserRepository
	jwtKey   []byte
}

func NewAuthService(userRepo repository.UserRepository, jwtKey []byte) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		jwtKey:   jwtKey,
	}
}

func (s *AuthService) Register(ctx context.Context, email, password string) error {
	// Validate input
	if email == "" || password == "" {
		return errors.ErrValidation
	}

	// Check if user already exists
	existingUser, err := s.userRepo.GetByEmail(ctx, email)
	if err == nil && existingUser != nil {
		return errors.ErrDuplicate
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.ErrInternal
	}

	// Create user
	user := &model.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return errors.ErrDatabase
	}

	return nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	// Validate input
	if email == "" || password == "" {
		return "", errors.ErrValidation
	}

	// Get user
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", errors.ErrUnauthorized
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.ErrUnauthorized
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(s.jwtKey)
	if err != nil {
		return "", errors.ErrInternal
	}

	return tokenString, nil
}
