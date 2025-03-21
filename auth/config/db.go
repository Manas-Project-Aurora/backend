package config

import (
	"log"

	"github.com/Manas-Project-Aurora/gavna/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnectToDB устанавливает соединение с базой данных SQLite.
// Функция также выполняет автоматическую миграцию для модели User,
// что обеспечивает создание (или обновление) таблицы в базе данных.
func ConnectToDB() (*gorm.DB, error) {
	// Открываем соединение с базой данных, в данном случае используем SQLite
	db, err := gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Автоматически применяем миграции для модели User
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Println("Ошибка миграции модели User:", err)
		return nil, err
	}

	return db, nil
}
