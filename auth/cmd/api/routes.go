package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/Manas-Project-Aurora/gavna/auth/handlers"
)

// RegisterRoutes регистрирует все маршруты для сервиса аутентификации.
// Здесь создается группа для версии API v1 и вложенная группа для auth.
func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// Группа для версии API v1
	apiV1 := router.Group("/v1")
	// Группа для аутентификации
	authGroup := apiV1.Group("/auth")
	// Регистрируем маршруты, связанные с аутентификацией (регистрация, вход, обновление токена, логаут)
	handlers.RegisterAuthRoutes(authGroup, db)
}
