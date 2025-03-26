package main

import (
	"log"

	"github.com/Manas-Project-Aurora/gavna/auth/cmd/api"
	"github.com/Manas-Project-Aurora/gavna/auth/config"
)

func main() {
	// Инициализация базы данных
	config.InitDB()

	// Настройка маршрутов
	router := api.SetupRoutes()

	// Запуск сервера
	log.Println("Starting auth service on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
