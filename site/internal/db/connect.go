package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(configPath string) (*gorm.DB, error) {
	config, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Name, config.Password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	AutomateMigrations(db)
	return db, nil
}
