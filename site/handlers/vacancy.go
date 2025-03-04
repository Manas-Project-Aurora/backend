package handlers

import (
	"net/http"
	"strconv"

	"github.com/Manas-Project-Aurora/gavna/internal/models"
	"github.com/Manas-Project-Aurora/gavna/site/repository"
	"github.com/Manas-Project-Aurora/gavna/site/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type VacancyHandler struct {
	Service services.VacancyService
}

func NewVacancyHandler(service services.VacancyService) *VacancyHandler {
	return &VacancyHandler{Service: service}
}

func RegisterVacancyRoutes(router *gin.RouterGroup, db *gorm.DB) {
	vacRepo := repository.NewVacancyRepository(db)
	vacService := services.NewVacancyService(vacRepo)
	vacHandler := NewVacancyHandler(vacService)

	vacRoutes := router.Group("/vacancies")
	{
		vacRoutes.GET("", vacHandler.GetVacancies)
		vacRoutes.GET("/:id", vacHandler.GetVacancyByID)
		vacRoutes.POST("", vacHandler.CreateVacancy)
		vacRoutes.PUT("/:id", vacHandler.UpdateVacancy)
		vacRoutes.DELETE("/:id", vacHandler.DeleteVacancy)
	}
}

func (h *VacancyHandler) GetVacancies(c *gin.Context) {
	take, _ := strconv.Atoi(c.DefaultQuery("take", "10"))
	skip, _ := strconv.Atoi(c.DefaultQuery("skip", "0"))

	vacs, total, err := h.Service.GetVacancies(take, skip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"vacanizations": vacs,
		"pagination": gin.H{
			"taken_count":   len(vacs),
			"skipped_count": skip,
			"total_count":   total,
		},
	})
}

func (h *VacancyHandler) GetVacancyByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("vacanization-id"))
	vac, err := h.Service.GetVacancyByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vacancy not found"})
		return
	}

	c.JSON(http.StatusOK, vac)
}

func (h *VacancyHandler) CreateVacancy(c *gin.Context) {
	var vac models.Vacancy
	if err := c.ShouldBindJSON(&vac); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateVacancy(&vac); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, vac)
}

func (h *VacancyHandler) UpdateVacancy(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("vacanization-id"))
	var vac models.Vacancy
	if err := c.ShouldBindJSON(&vac); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.UpdateVacancy(uint(id), &vac); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *VacancyHandler) DeleteVacancy(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("vacanization-id"))
	if err := h.Service.DeleteVacancy(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
