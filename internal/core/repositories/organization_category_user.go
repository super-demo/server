package repositories

import (
	"server/internal/core/models"

	"gorm.io/gorm"
)

type OrganizationCategoryUserRepository interface {
	BeginLog() (OrganizationCategoryUserRepository, error)
	Commit() error
	Rollback() error
	CreateOrganizationCategoryUser(organizationCategoryUser *models.OrganizationCategoryUser) (*models.OrganizationCategoryUser, error)
	CheckOrganizationCategoryUserExistsById(id int) (bool, error)
}

type organizationCategoryUserRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewOrganizationCategoryUserRepository(db *gorm.DB) OrganizationCategoryUserRepository {
	return &organizationCategoryUserRepository{db: db}
}

func (r *organizationCategoryUserRepository) BeginLog() (OrganizationCategoryUserRepository, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &organizationCategoryUserRepository{db: tx, tx: tx}, nil
}

func (r *organizationCategoryUserRepository) Commit() error {
	return r.tx.Commit().Error
}

func (r *organizationCategoryUserRepository) Rollback() error {
	return r.tx.Rollback().Error
}

func (r *organizationCategoryUserRepository) CreateOrganizationCategoryUser(organizationCategoryUser *models.OrganizationCategoryUser) (*models.OrganizationCategoryUser, error) {
	if err := r.db.Create(organizationCategoryUser).Error; err != nil {
		return nil, err
	}

	return organizationCategoryUser, nil
}

func (r *organizationCategoryUserRepository) CheckOrganizationCategoryUserExistsById(id int) (bool, error) {
	var organizationCategoryUser models.OrganizationCategoryUser

	err := r.db.Where("user_id = ?", id).First(&organizationCategoryUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
	}

	return true, nil
}
