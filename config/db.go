package config

import (
	"fmt"
	"log"

	"github.com/sandeep-jaiswar/cms-backend/internal/models"
	"github.com/sandeep-jaiswar/cms-backend/internal/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		AppConfig.Database.Host,
		AppConfig.Database.Port,
		AppConfig.Database.User,
		AppConfig.Database.Password,
		AppConfig.Database.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db

	if err := DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.Product{},
		&models.Category{},
		&models.Tag{},
		&models.Media{},
		&models.Customer{},
		&models.Order{},
		&models.OrderItem{},
		&models.Inventory{},
		&models.Review{},
		&models.Discount{},
		&models.AuditLog{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	repositories.InitRepositories(DB)
}
