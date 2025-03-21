package main

import (
	"fmt"
	"log"

	"github.com/Manas-Project-Aurora/gavna/auth/config"

	"github.com/Manas-Project-Aurora/gavna/auth/cmd/api"

	"github.com/gin-gonic/gin"
)

// main — точка входа в приложение.
// Здесь происходит инициализация подключения к базе данных,
// создание маршрутизатора Gin, регистрация маршрутов API и запуск сервера.
func main() {
	// Устанавливаем соединение с базой данных
	db, err := config.ConnectToDB()
	if err != nil {
		panic(err)
	}

	// Инициализируем маршрутизатор Gin
	router := gin.Default()

	// Регистрируем маршруты API, передавая объект router и подключение к БД
	api.RegisterRoutes(router, db)

	log.Println("Server running on 8080")
	fmt.Println("http://localhost:8080/v1/")
	// Запускаем сервер на порту 8080
	router.Run(":8080")
}
