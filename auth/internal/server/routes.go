package server

import (
	"github.com/Manas-Project-Aurora/gavna/auth/internal/handlers"
	"github.com/Manas-Project-Aurora/gavna/auth/internal/repository"
	"github.com/Manas-Project-Aurora/gavna/auth/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Инициализация репозитория, сервиса и хендлера
	authRepo := repository.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	// Группы роутов
	apiV1 := router.Group("/v1")
	authGroup := apiV1.Group("/aurora-auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/token", authHandler.RefreshToken)
		authGroup.POST("/logout", authHandler.Logout)
	}

	return router
}
