package repositories

import (
	"server/internal/core/models"

	"gorm.io/gorm"
)

type OrganizationCategoryServiceRepository interface {
	BeginLog() (OrganizationCategoryServiceRepository, error)
	Commit() error
	Rollback() error
	CreateOrganizationCategoryService(organizationCategoryService *models.OrganizationCategoryService) (*models.OrganizationCategoryService, error)
	CheckOrganizationCategoryServiceExistsById(id int) (bool, error)
}

type organizationCategoryServiceRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewOrganizationCategoryServiceRepository(db *gorm.DB) OrganizationCategoryServiceRepository {
	return &organizationCategoryServiceRepository{db: db}
}

func (r *organizationCategoryServiceRepository) BeginLog() (OrganizationCategoryServiceRepository, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &organizationCategoryServiceRepository{db: tx, tx: tx}, nil
}

func (r *organizationCategoryServiceRepository) Commit() error {
	return r.tx.Commit().Error
}

func (r *organizationCategoryServiceRepository) Rollback() error {
	return r.tx.Rollback().Error
}

func (r *organizationCategoryServiceRepository) CreateOrganizationCategoryService(organizationCategoryService *models.OrganizationCategoryService) (*models.OrganizationCategoryService, error) {
	if err := r.db.Create(organizationCategoryService).Error; err != nil {
		return nil, err
	}

	return organizationCategoryService, nil
}

func (r *organizationCategoryServiceRepository) CheckOrganizationCategoryServiceExistsById(id int) (bool, error) {
	var organizationCategoryService models.OrganizationCategoryService

	err := r.db.Where("organization_service_id = ?", id).First(&organizationCategoryService).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
	}

	return true, nil
}
