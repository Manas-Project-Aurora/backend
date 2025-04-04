package repository

import (
	"github.com/Manas-Project-Aurora/backend/internal/models"
	"gorm.io/gorm"
)

type VacancyRepository interface {
	GetAllVacancies(take, skip int) ([]models.Vacancy, int64, error)
	GetVacancyByID(id uint) (models.Vacancy, error)
	CreateVacancy(vac *models.Vacancy) error
	UpdateVacancy(vac *models.Vacancy) error
	DeleteVacancy(id uint) error
}

type vacancyRepository struct {
	db *gorm.DB
}

func NewVacancyRepository(db *gorm.DB) VacancyRepository {
	return &vacancyRepository{db: db}
}

func (r *vacancyRepository) GetAllVacancies(take, skip int) ([]models.Vacancy, int64, error) {
	var vacs []models.Vacancy
	var total int64

	err := r.db.Model(&models.Vacancy{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("Org").Offset(skip).Limit(take).Find(&vacs).Error
	return vacs, total, err
}

func (r *vacancyRepository) GetVacancyByID(id uint) (models.Vacancy, error) {
	var vac models.Vacancy
	err := r.db.Preload("Org").First(&vac, id).Error
	return vac, err
}

func (r *vacancyRepository) CreateVacancy(vac *models.Vacancy) error {
	return r.db.Create(vac).Error
}

func (r *vacancyRepository) UpdateVacancy(vac *models.Vacancy) error {
	return r.db.Save(vac).Error
}

func (r *vacancyRepository) DeleteVacancy(id uint) error {
	return r.db.Delete(&models.Vacancy{}, id).Error
}
