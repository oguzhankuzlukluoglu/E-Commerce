package database

import (
	"fmt"
	"log"

	"github.com/oguzhan/e-commerce/pkg/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	models := []interface{}{
		&models.User{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
		&models.Payment{},
	}

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return fmt.Errorf("failed to migrate model: %v", err)
		}
	}

	log.Println("Successfully migrated database schema")
	return nil
}
