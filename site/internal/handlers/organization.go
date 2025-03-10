package handlers

import (
	"net/http"
	"strconv"

	"github.com/Manas-Project-Aurora/gavna/internal/models"
	"github.com/Manas-Project-Aurora/gavna/site/internal/repository"
	"github.com/Manas-Project-Aurora/gavna/site/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrganizationHandler struct {
	Service services.OrganizationService
}

func NewOrganizationHandler(service services.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{Service: service}
}

func RegisterOrganizationRoutes(router *gin.RouterGroup, db *gorm.DB) {
	orgRepo := repository.NewOrganizationRepository(db)
	orgService := services.NewOrganizationService(orgRepo)
	orgHandler := NewOrganizationHandler(orgService)

	orgRoutes := router.Group("/organizations")
	{
		orgRoutes.GET("", orgHandler.GetOrganizations)
		orgRoutes.GET("/:organization-id", orgHandler.GetOrganizationByID)
		orgRoutes.POST("", orgHandler.CreateOrganization)
		orgRoutes.PUT("/:organization-id", orgHandler.UpdateOrganization)
		orgRoutes.DELETE("/:organization-id", orgHandler.DeleteOrganization)
	}
}

func (h *OrganizationHandler) GetOrganizations(c *gin.Context) {
	take, _ := strconv.Atoi(c.DefaultQuery("take", "10"))
	skip, _ := strconv.Atoi(c.DefaultQuery("skip", "0"))

	orgs, total, err := h.Service.GetOrganizations(take, skip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"organizations": orgs,
		"pagination": gin.H{
			"taken_count":   len(orgs),
			"skipped_count": skip,
			"total_count":   total,
		},
	})
}

func (h *OrganizationHandler) GetOrganizationByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("organization-id"))
	org, err := h.Service.GetOrganizationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
		return
	}

	c.JSON(http.StatusOK, org)
}

func (h *OrganizationHandler) CreateOrganization(c *gin.Context) {
	var org models.Organization
	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateOrganization(&org); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, org)
}

func (h *OrganizationHandler) UpdateOrganization(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("organization-id"))
	var org models.Organization
	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.UpdateOrganization(uint(id), &org); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *OrganizationHandler) DeleteOrganization(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("organization-id"))
	if err := h.Service.DeleteOrganization(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
