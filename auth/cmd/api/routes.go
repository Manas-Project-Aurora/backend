package api

import (
	"github.com/Manas-Project-Aurora/gavna/auth/config"
	"github.com/Manas-Project-Aurora/gavna/auth/handlers"
	"github.com/Manas-Project-Aurora/gavna/auth/repository"
	"github.com/Manas-Project-Aurora/gavna/auth/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Инициализация репозитория, сервиса и хендлера
	authRepo := repository.NewAuthRepository(config.DB)
	authService := services.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	// Группы роутов
	apiV1 := router.Group("/v1")
	authGroup := apiV1.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/token", authHandler.RefreshToken)
		authGroup.POST("/logout", authHandler.Logout)
	}

	return router
}
