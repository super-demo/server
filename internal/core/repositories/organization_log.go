package repositories

import (
	"server/internal/core/models"

	"gorm.io/gorm"
)

type OrganizationLogRepository interface {
	CreateOrganizationLog(organizationLog *models.OrganizationLog) (*models.OrganizationLog, error)
}

type organizationLogRepository struct {
	db *gorm.DB
}

func NewOrganizationLogRepository(db *gorm.DB) OrganizationLogRepository {
	return &organizationLogRepository{db}
}

func (r *organizationLogRepository) CreateOrganizationLog(organizationLog *models.OrganizationLog) (*models.OrganizationLog, error) {
	if err := r.db.Create(organizationLog).Error; err != nil {
		return nil, err
	}

	return organizationLog, nil
}
