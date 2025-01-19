package repositories

import (
	"server/internal/core/models"

	"gorm.io/gorm"
)

type OrganizationServiceRepository interface {
	BeginLog() (OrganizationServiceRepository, error)
	Commit() error
	Rollback() error
	CreateOrganizationService(organizationService *models.OrganizationService) (*models.OrganizationService, error)
	CheckOrganizationServiceExists(id, serviceId int) (bool, error)
	CheckOrganizationServiceExistsByName(name string) (bool, error)
	DeleteOrganizationService(organizationService *models.OrganizationService) error
}

type organizationServiceRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewOrganizationServiceRepository(db *gorm.DB) OrganizationServiceRepository {
	return &organizationServiceRepository{db: db}
}

func (r *organizationServiceRepository) BeginLog() (OrganizationServiceRepository, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &organizationServiceRepository{db: tx, tx: tx}, nil
}

func (r *organizationServiceRepository) Commit() error {
	return r.tx.Commit().Error
}

func (r *organizationServiceRepository) Rollback() error {
	return r.tx.Rollback().Error
}

func (r *organizationServiceRepository) CreateOrganizationService(organizationService *models.OrganizationService) (*models.OrganizationService, error) {
	if err := r.db.Create(organizationService).Error; err != nil {
		return nil, err
	}

	return organizationService, nil
}

func (r *organizationServiceRepository) CheckOrganizationServiceExists(id, serviceId int) (bool, error) {
	var organizationService models.OrganizationService

	if err := r.db.Where("organization_id = ?", id).Where("organization_service_id = ?", serviceId).First(&organizationService).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *organizationServiceRepository) CheckOrganizationServiceExistsByName(name string) (bool, error) {
	var organizationService models.OrganizationService
	if err := r.db.Where("slug = ?", name).First(&organizationService).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (r *organizationServiceRepository) DeleteOrganizationService(organizationService *models.OrganizationService) error {
	if err := r.db.Where("organization_service_id = ?", organizationService.OrganizationServiceId).Delete(organizationService).Error; err != nil {
		return err
	}

	return nil
}
