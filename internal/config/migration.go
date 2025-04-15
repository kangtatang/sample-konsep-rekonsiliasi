// File: config/migration.go
package config

import (
	"go-big-internal/models"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.InternalTransaction{},
		&models.ReconciliationResult{},
	)

	if err != nil {
		log.Println("Migration error:", err)
		return err
	}

	log.Println("Database migration successful")
	return nil
}
