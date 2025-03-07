package config

import (
	"github.com/Manas-Project-Aurora/gavna/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("organizations.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Organization{})
	return db, nil
}
