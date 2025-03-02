package main

import (
	"github.com/Manas-Project-Aurora/gavna/site/handlers"
	"github.com/Manas-Project-Aurora/gavna/site/repository"
	"github.com/Manas-Project-Aurora/gavna/site/services"
  "github.com/Manas-Project-Aurora/gavna/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("organizations.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto-migrate the schema
	db.AutoMigrate(&models.Organization{})

	repo := repository.NewOrganizationRepository(db)
	service := services.NewOrganizationService(repo)
	handler := handlers.NewOrganizationHandler(service)

	r := gin.Default()

	r.GET("/v1/organizations", handler.GetOrganizations)
	r.GET("/v1/organizations/:organization-id", handler.GetOrganizationByID)
	r.POST("/v1/organizations", handler.CreateOrganization)
	r.PUT("/v1/organizations/:organization-id", handler.UpdateOrganization)
	r.DELETE("/v1/organizations/:organization-id", handler.DeleteOrganization)

	r.Run(":8080")
}

