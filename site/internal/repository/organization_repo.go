package repository

import (
	"github.com/Manas-Project-Aurora/backend/internal/models"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	GetAllOrganizations(take, skip int) ([]models.Organization, int64, error)
	GetOrganizationByID(id uint) (models.Organization, error)
	CreateOrganization(org *models.Organization) error
	UpdateOrganization(org *models.Organization) error
	DeleteOrganization(id uint) error
}

type organizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationRepository{db: db}
}

func (r *organizationRepository) GetAllOrganizations(take, skip int) ([]models.Organization, int64, error) {
	var orgs []models.Organization
	var total int64

	err := r.db.Model(&models.Organization{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Offset(skip).Limit(take).Find(&orgs).Error
	return orgs, total, err
}

func (r *organizationRepository) GetOrganizationByID(id uint) (models.Organization, error) {
	var org models.Organization
	err := r.db.First(&org, id).Error
	return org, err
}

func (r *organizationRepository) CreateOrganization(org *models.Organization) error {
	return r.db.Create(org).Error
}

func (r *organizationRepository) UpdateOrganization(org *models.Organization) error {
	return r.db.Save(org).Error
}

func (r *organizationRepository) DeleteOrganization(id uint) error {
	return r.db.Delete(&models.Organization{}, id).Error
}
