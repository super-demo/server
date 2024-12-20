package repositories

import (
	"server/internal/core/models"

	"gorm.io/gorm"
)

type OrganizationRepository interface {
	CreateOrganization(organization *models.Organization) (*models.Organization, error)
	GetOrganizationListByUserId(id int) (*[]models.Organization, error)
}

type organizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationRepository{db}
}

func (r *organizationRepository) CreateOrganization(organization *models.Organization) (*models.Organization, error) {
	if err := r.db.Create(organization).Error; err != nil {
		return organization, err
	}

	return organization, nil

}

func (r *organizationRepository) GetOrganizationListByUserId(id int) (*[]models.Organization, error) {
	organization := new([]models.Organization)

	if err := r.db.Where("created_by = ?", id).Find(organization).Error; err != nil {
		return nil, err
	}

	return organization, nil
}
