package auth

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/oguzhan/e-commerce/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db:     db,
		logger: log.Default(),
	}
}

func (s *Service) Register(user *models.User) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	// Create a new user with the hashed password
	newUser := &models.User{
		Email:     user.Email,
		Password:  string(hashedPassword),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      "user",
	}

	// Save the user to the database
	if err := s.db.Create(newUser).Error; err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	// Copy the created user's data back to the input user
	user.ID = newUser.ID
	user.CreatedAt = newUser.CreatedAt
	user.UpdatedAt = newUser.UpdatedAt
	user.DeletedAt = newUser.DeletedAt
	user.Role = newUser.Role
	user.LastLogin = newUser.LastLogin

	return nil
}

func (s *Service) Login(email, password string) (*models.User, error) {
	// Find the user by email
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to retrieve user: %v", err)
	}

	// Compare the provided password with the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		s.logger.Printf("Password comparison failed: %v", err)
		return nil, errors.New("invalid password")
	}

	// Update last login time
	now := time.Now()
	user.LastLogin = now
	if err := s.db.Model(&user).Update("last_login", now).Error; err != nil {
		return nil, fmt.Errorf("failed to update last login time: %v", err)
	}

	return &user, nil
}

func (s *Service) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Service) generateToken(user *models.User) (string, time.Time, error) {
	expiresAt := time.Now().Add(24 * time.Hour)
	claims := &models.Claims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expiresAt, nil
}

func (s *Service) ValidateToken(tokenString string) (*models.Claims, error) {
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
