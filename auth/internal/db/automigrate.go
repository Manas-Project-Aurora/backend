package db

import (
	"github.com/Manas-Project-Aurora/backend/auth/internal/models"
	in "github.com/Manas-Project-Aurora/backend/internal/models"
	"gorm.io/gorm"
	"log"
	"reflect"
)

func AutomateMigrations(db *gorm.DB) {
	modelsToMigrate := []interface{}{
		&models.RefreshToken{},
		&in.User{},
	}

	for _, model := range modelsToMigrate {
		if err := db.AutoMigrate(model); err != nil {
			log.Fatalf("Auto migration failed for %v: %v", reflect.TypeOf(model), err)
		}
	}
}
