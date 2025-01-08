package repositories

import (
	"server/internal/core/models"

	"gorm.io/gorm"
)

type OrganizationRepository interface {
	BeginTransaction() (OrganizationRepository, error)
	Commit() error
	Rollback() error
	CreateOrganization(organization *models.Organization) (*models.Organization, error)
	DeleteOrganization(id int) error
	GetOrganizationById(id int) (*models.Organization, error)
	GetOrganizationListByUserId(id int) (*[]models.Organization, error)
}

type organizationRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationRepository{db: db}
}

func (r *organizationRepository) BeginTransaction() (OrganizationRepository, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &organizationRepository{db: tx, tx: tx}, nil
}

func (r *organizationRepository) Commit() error {
	return r.tx.Commit().Error
}

func (r *organizationRepository) Rollback() error {
	return r.tx.Rollback().Error
}

func (r *organizationRepository) CreateOrganization(organization *models.Organization) (*models.Organization, error) {
	if err := r.db.Create(organization).Error; err != nil {
		return organization, err
	}

	return organization, nil

}

func (r *organizationRepository) DeleteOrganization(id int) error {
	if err := r.db.Where("organization_id = ?", id).Delete(&models.Organization{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *organizationRepository) GetOrganizationById(id int) (*models.Organization, error) {
	organization := new(models.Organization)

	if err := r.db.First(organization, id).Error; err != nil {
		return nil, err
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
