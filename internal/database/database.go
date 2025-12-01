package database

import (
	"fmt"
	"log"
	"os"

	"github.com/emuthianimbithi/pos-service/internal/config"
	"github.com/emuthianimbithi/pos-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Initialize(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode,
	)

	// Set logger level based on environment
	logLevel := logger.Info
	if cfg.Environment == "production" {
		logLevel = logger.Error
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto-migrate audit model and mpesa models
	// In production (Cloud Run), you might want to run migrations separately
	if os.Getenv("SKIP_MIGRATIONS") != "true" {
		if err := db.AutoMigrate(
			&models.AuditLog{},
			&models.MpesaConfig{},
			&models.MpesaTransaction{},
		); err != nil {
			log.Printf("Warning: Failed to migrate database: %v", err)
		}
	}

	log.Println("Database connection established")
	return db, nil
}
