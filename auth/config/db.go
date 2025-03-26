package config

import (
	"log"

	authmodels "github.com/Manas-Project-Aurora/gavna/auth/models"
	internalmodels "github.com/Manas-Project-Aurora/gavna/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Миграция таблиц
	err = DB.AutoMigrate(&internalmodels.User{}, &authmodels.RefreshToken{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database connected and migrated successfully")
}
