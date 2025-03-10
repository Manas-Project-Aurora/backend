package server

import (
	"github.com/Manas-Project-Aurora/gavna/site/internal/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	apiV1 := router.Group("/v1")

	handlers.RegisterOrganizationRoutes(apiV1, db)
	handlers.RegisterVacancyRoutes(apiV1, db)
}
