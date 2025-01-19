package repositories

import (
	"errors"
	"server/infrastructure/app"
	"server/internal/core/models"

	"gorm.io/gorm"
)

type OrganizationRepository interface {
	BeginLog() (OrganizationRepository, error)
	Commit() error
	Rollback() error
	CreateOrganization(organization *models.Organization) (*models.Organization, error)
	CheckOrganizationExistsByName(name string) (bool, error)
	GetOrganizationById(id int) (*models.Organization, error)
	GetOrganizationListByUserId(id int) (*[]models.Organization, error)
	DeleteOrganization(organization *models.Organization) error
}

type organizationRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationRepository{db: db}
}

func (r *organizationRepository) BeginLog() (OrganizationRepository, error) {
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
		return nil, err
	}

	return organization, nil
}

func (r *organizationRepository) CheckOrganizationExistsByName(name string) (bool, error) {
	var organization models.Organization

	err := r.db.Where("name = ?", name).First(&organization).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, app.ErrOrganizationNameExists
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

func (r *organizationRepository) DeleteOrganization(organization *models.Organization) error {
	if err := r.db.Where("organization_id = ?", organization.OrganizationId).Delete(organization).Error; err != nil {
		return err
	}

	return nil
}
