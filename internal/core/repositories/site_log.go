package repositories

import (
	"server/internal/core/models"

	"gorm.io/gorm"
)

type SiteLogRepository interface {
	BeginLog() (SiteLogRepository, error)
	Commit() error
	Rollback() error
	CreateSiteLog(siteLog *models.SiteLog) (*models.SiteLog, error)
}

type siteLogRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewSiteLogRepository(db *gorm.DB) SiteLogRepository {
	return &siteLogRepository{db: db}
}

func (r *siteLogRepository) BeginLog() (SiteLogRepository, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &siteLogRepository{db: r.db, tx: tx}, nil
}

func (r *siteLogRepository) Commit() error {
	return r.tx.Commit().Error
}

func (r *siteLogRepository) Rollback() error {
	return r.tx.Rollback().Error
}

func (r *siteLogRepository) CreateSiteLog(siteLog *models.SiteLog) (*models.SiteLog, error) {
	if err := r.db.Create(siteLog).Error; err != nil {
		return nil, err
	}

	return siteLog, nil
}
