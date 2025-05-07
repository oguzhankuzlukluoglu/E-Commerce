package service

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/oguzhan/e-commerce/internal/user/domain"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo      domain.UserRepository
	jwtSecret string
}

func NewUserService(repo domain.UserRepository, jwtSecret string) domain.UserService {
	return &userService{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *userService) Register(user *domain.User) error {
	// Check if user already exists
	existingUser, err := s.repo.GetByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("user already exists")
	}

	// Set default role if not provided
	if user.Role == "" {
		user.Role = "user"
	}

	return s.repo.Create(user)
}

func (s *userService) Login(email, password string) (string, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *userService) GetUserByID(id uint) (*domain.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) UpdateUser(user *domain.User) error {
	existingUser, err := s.repo.GetByID(user.ID)
	if err != nil {
		return err
	}

	// Update only allowed fields
	existingUser.Email = user.Email
	existingUser.Name = user.Name
	existingUser.Role = user.Role

	return s.repo.Update(existingUser)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
