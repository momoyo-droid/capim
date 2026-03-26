package postgres

import (
	"fmt"

	"github.com/momoyo-droid/capim/api/internal/config"
	"github.com/momoyo-droid/capim/api/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("create database connection error: %w", err)
	}
	// Create tables if they do not exist
	db.AutoMigrate(&repository.Seller{}, &repository.Owner{})

	return db, nil
}
