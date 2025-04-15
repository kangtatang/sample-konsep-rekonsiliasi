package config

import (
	"go-big-external/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.ExternalTransaction{})
}
