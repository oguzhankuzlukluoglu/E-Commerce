package service

import (
	"context"

	"github.com/oguzhan/e-commerce/internal/model"
	"github.com/oguzhan/e-commerce/internal/repository"
	"github.com/oguzhan/e-commerce/pkg/errors"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetProfile(ctx context.Context, userID uint) (*model.User, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, errors.ErrNotFound
	}
	return user, nil
}

func (s *UserService) UpdateProfile(ctx context.Context, userID uint, updateData map[string]interface{}) error {
	// Get user
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return errors.ErrNotFound
	}

	// Update fields
	if name, ok := updateData["name"].(string); ok {
		user.Name = name
	}

	// Save changes
	if err := s.userRepo.Update(ctx, user); err != nil {
		return errors.ErrDatabase
	}

	return nil
}
