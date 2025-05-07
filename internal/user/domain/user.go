package domain

import (
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	Create(user *User) error
	GetByID(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(user *User) error
	Delete(id uint) error
}

type UserService interface {
	Register(user *User) error
	Login(email, password string) (string, error)
	GetUserByID(id uint) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id uint) error
}
