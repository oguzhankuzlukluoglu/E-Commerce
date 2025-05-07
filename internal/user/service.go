package user

import (
	"errors"
	"time"

	"github.com/oguzhan/e-commerce/pkg/models"
	"golang.org/x/crypto/bcrypt"
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
	var existingUser models.User
	if err := s.db.First(&existingUser, id).Error; err != nil {
		return err
	}

	return s.db.Model(&existingUser).Updates(user).Error
}

func (s *Service) DeleteUser(id uint) error {
	return s.db.Delete(&models.User{}, id).Error
}

func (s *Service) UpdateLastLogin(id uint) error {
	return s.db.Model(&models.User{}).Where("id = ?", id).Update("last_login", time.Now()).Error
}

func (s *Service) ListUsers(page, limit int) ([]models.User, int64, error) {
	var users []models.User
	var total int64
	if err := s.db.Limit(limit).Offset((page - 1) * limit).Find(&users).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (s *Service) ChangePassword(id uint, oldPassword, newPassword string) error {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return errors.New("invalid old password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.db.Model(&user).Update("password", string(hashedPassword)).Error
}

func (s *Service) DeactivateUser(id uint) error {
	return s.db.Model(&models.User{}).Where("id = ?", id).Update("is_active", false).Error
}

func (s *Service) ActivateUser(id uint) error {
	return s.db.Model(&models.User{}).Where("id = ?", id).Update("is_active", true).Error
}

func (s *Service) UpdateUserRole(id uint, role string) error {
	return s.db.Model(&models.User{}).Where("id = ?", id).Update("role", role).Error
}

func (s *Service) ResetPassword(id uint, newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return s.db.Model(&models.User{}).Where("id = ?", id).Update("password", string(hashedPassword)).Error
}

// Address Management
func (s *Service) CreateAddress(address *models.Address) error {
	if address.IsDefault {
		// Set all other addresses to non-default
		if err := s.db.Model(&models.Address{}).Where("user_id = ?", address.UserID).Update("is_default", false).Error; err != nil {
			return err
		}
	}
	return s.db.Create(address).Error
}

func (s *Service) GetAddresses(userID uint) ([]models.Address, error) {
	var addresses []models.Address
	if err := s.db.Where("user_id = ?", userID).Find(&addresses).Error; err != nil {
		return nil, err
	}
	return addresses, nil
}

func (s *Service) UpdateAddress(id uint, address *models.Address) error {
	if address.IsDefault {
		// Set all other addresses to non-default
		if err := s.db.Model(&models.Address{}).Where("user_id = ? AND id != ?", address.UserID, id).Update("is_default", false).Error; err != nil {
			return err
		}
	}
	return s.db.Model(&models.Address{}).Where("id = ?", id).Updates(address).Error
}

func (s *Service) DeleteAddress(id uint) error {
	return s.db.Delete(&models.Address{}, id).Error
}

// Contact Management
func (s *Service) CreateContact(contact *models.Contact) error {
	if contact.IsDefault {
		// Set all other contacts to non-default
		if err := s.db.Model(&models.Contact{}).Where("user_id = ?", contact.UserID).Update("is_default", false).Error; err != nil {
			return err
		}
	}
	return s.db.Create(contact).Error
}

func (s *Service) GetContacts(userID uint) ([]models.Contact, error) {
	var contacts []models.Contact
	if err := s.db.Where("user_id = ?", userID).Find(&contacts).Error; err != nil {
		return nil, err
	}
	return contacts, nil
}

func (s *Service) UpdateContact(id uint, contact *models.Contact) error {
	if contact.IsDefault {
		// Set all other contacts to non-default
		if err := s.db.Model(&models.Contact{}).Where("user_id = ? AND id != ?", contact.UserID, id).Update("is_default", false).Error; err != nil {
			return err
		}
	}
	return s.db.Model(&models.Contact{}).Where("id = ?", id).Updates(contact).Error
}

func (s *Service) DeleteContact(id uint) error {
	return s.db.Delete(&models.Contact{}, id).Error
}
