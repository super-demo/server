package repositories

import (
	"server/internal/core/models"

	"gorm.io/gorm"
)

type OrganizationCategoryRepository interface {
	BeginLog() (OrganizationCategoryRepository, error)
	Commit() error
	Rollback() error
	CreateOrganizationCategory(organizationCategory *models.OrganizationCategory) (*models.OrganizationCategory, error)
	CheckOrganizationCategoryExistsByName(name string) (bool, error)
}

type organizationCategoryRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewOrganizationCategoryRepository(db *gorm.DB) OrganizationCategoryRepository {
	return &organizationCategoryRepository{db: db}
}

func (r *organizationCategoryRepository) BeginLog() (OrganizationCategoryRepository, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &organizationCategoryRepository{db: tx, tx: tx}, nil
}

func (r *organizationCategoryRepository) Commit() error {
	return r.tx.Commit().Error
}

func (r *organizationCategoryRepository) Rollback() error {
	return r.tx.Rollback().Error
}

func (r *organizationCategoryRepository) CreateOrganizationCategory(organizationCategory *models.OrganizationCategory) (*models.OrganizationCategory, error) {
	if err := r.db.Create(organizationCategory).Error; err != nil {
		return nil, err
	}

	return organizationCategory, nil
}

func (r *organizationCategoryRepository) CheckOrganizationCategoryExistsByName(name string) (bool, error) {
	var organizationCategory models.OrganizationCategory
	if err := r.db.Where("name = ?", name).First(&organizationCategory).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
