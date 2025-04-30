package user

import (
	"errors"
	"time"

	"github.com/oguzhan/e-commerce/pkg/models"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateUser(user *models.User) error {
	return s.db.Create(user).Error
}

func (s *Service) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (s *Service) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (s *Service) UpdateUser(id uint, user *models.User) error {
	existingUser, err := s.GetUserByID(id)
	if err != nil {
		return err
	}

	// Update only non-zero fields
	if user.Email != "" {
		existingUser.Email = user.Email
	}
	if user.FirstName != "" {
		existingUser.FirstName = user.FirstName
	}
	if user.LastName != "" {
		existingUser.LastName = user.LastName
	}
	if user.Role != "" {
		existingUser.Role = user.Role
	}

	return s.db.Save(existingUser).Error
}

func (s *Service) DeleteUser(id uint) error {
	return s.db.Delete(&models.User{}, id).Error
}

func (s *Service) UpdateLastLogin(id uint) error {
	return s.db.Model(&models.User{}).Where("id = ?", id).Update("last_login", time.Now()).Error
}
