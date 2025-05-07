package models

import (
	"time"

	"gorm.io/gorm"
)

type AddressType string

const (
	AddressTypeHome AddressType = "home"
	AddressTypeWork AddressType = "work"
)

type ContactType string

const (
	ContactTypeHome ContactType = "home"
	ContactTypeWork ContactType = "work"
)

type User struct {
	gorm.Model
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Password  string    `json:"password" gorm:"not null"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Role      string    `json:"role" gorm:"default:user"`
	LastLogin time.Time `json:"last_login"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	Addresses []Address `json:"addresses"`
	Contacts  []Contact `json:"contacts"`
}

type Address struct {
	gorm.Model
	UserID      uint        `json:"user_id" gorm:"not null"`
	Type        AddressType `json:"type" gorm:"type:varchar(10);not null"`
	Title       string      `json:"title" gorm:"not null"`
	AddressLine string      `json:"address_line" gorm:"not null"`
	City        string      `json:"city" gorm:"not null"`
	State       string      `json:"state"`
	Country     string      `json:"country" gorm:"not null"`
	PostalCode  string      `json:"postal_code" gorm:"not null"`
	IsDefault   bool        `json:"is_default" gorm:"default:false"`
}

type Contact struct {
	gorm.Model
	UserID      uint        `json:"user_id" gorm:"not null"`
	Type        ContactType `json:"type" gorm:"type:varchar(10);not null"`
	Title       string      `json:"title" gorm:"not null"`
	PhoneNumber string      `json:"phone_number" gorm:"not null"`
	IsDefault   bool        `json:"is_default" gorm:"default:false"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	return nil
}
