package services

import (
	"github.com/Manas-Project-Aurora/backend/internal/models"
	"github.com/Manas-Project-Aurora/backend/site/internal/repository"
)

type OrganizationService interface {
	GetOrganizations(take, skip int) ([]models.Organization, int64, error)
	GetOrganizationByID(id uint) (models.Organization, error)
	CreateOrganization(org *models.Organization) error
	UpdateOrganization(id uint, org *models.Organization) error
	DeleteOrganization(id uint) error
}

type organizationService struct {
	repo repository.OrganizationRepository
}

func NewOrganizationService(repo repository.OrganizationRepository) OrganizationService {
	return &organizationService{repo: repo}
}

func (s *organizationService) GetOrganizations(take, skip int) ([]models.Organization, int64, error) {
	return s.repo.GetAllOrganizations(take, skip)
}

func (s *organizationService) GetOrganizationByID(id uint) (models.Organization, error) {
	return s.repo.GetOrganizationByID(id)
}

func (s *organizationService) CreateOrganization(org *models.Organization) error {
	return s.repo.CreateOrganization(org)
}

func (s *organizationService) UpdateOrganization(id uint, org *models.Organization) error {
	existingOrg, err := s.repo.GetOrganizationByID(id)
	if err != nil {
		return err
	}
	existingOrg = *org

	return s.repo.UpdateOrganization(&existingOrg)
}

func (s *organizationService) DeleteOrganization(id uint) error {
	return s.repo.DeleteOrganization(id)
}
