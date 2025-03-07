package services

import (
	"github.com/Manas-Project-Aurora/gavna/internal/models"
	"github.com/Manas-Project-Aurora/gavna/site/repository"
)

type VacancyService interface {
	GetVacancies(take, skip int) ([]models.Vacancy, int64, error)
	GetVacancyByID(id uint) (models.Vacancy, error)
	CreateVacancy(vac *models.Vacancy) error
	UpdateVacancy(id uint, vac *models.Vacancy) error
	DeleteVacancy(id uint) error
}

type vacancyService struct {
	repo repository.VacancyRepository
}

func NewVacancyService(repo repository.VacancyRepository) VacancyService {
	return &vacancyService{repo: repo}
}

func (s *vacancyService) GetVacancies(take, skip int) ([]models.Vacancy, int64, error) {
	return s.repo.GetAllVacancies(take, skip)
}

func (s *vacancyService) GetVacancyByID(id uint) (models.Vacancy, error) {
	return s.repo.GetVacancyByID(id)
}

func (s *vacancyService) CreateVacancy(vac *models.Vacancy) error {
	return s.repo.CreateVacancy(vac)
}

func (s *vacancyService) UpdateVacancy(id uint, vac *models.Vacancy) error {
	existingVac, err := s.repo.GetVacancyByID(id)
	if err != nil {
		return err
	}

	existingVac = *vac

	return s.repo.UpdateVacancy(&existingVac)
}

func (s *vacancyService) DeleteVacancy(id uint) error {
	return s.repo.DeleteVacancy(id)
}
