package repositories

import (
	"server/internal/core/models"

	"gorm.io/gorm"
)

type SiteUserRepository interface {
	BeginLog() (SiteUserRepository, error)
	Commit() error
	Rollback() error
	CreateSiteUser(siteUser *models.SiteUser) (*models.SiteUser, error)
	CheckSiteUserExistsBySiteIdAndUserId(siteId, userId int) (bool, error)
}

type siteUserRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewSiteUserRepository(db *gorm.DB) SiteUserRepository {
	return &siteUserRepository{db: db}
}

func (r *siteUserRepository) BeginLog() (SiteUserRepository, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &siteUserRepository{db: r.db, tx: tx}, nil
}

func (r *siteUserRepository) Commit() error {
	return r.tx.Commit().Error
}

func (r *siteUserRepository) Rollback() error {
	return r.tx.Rollback().Error
}

func (r *siteUserRepository) CreateSiteUser(siteUser *models.SiteUser) (*models.SiteUser, error) {
	if err := r.db.Create(siteUser).Error; err != nil {
		return nil, err
	}

	return siteUser, nil
}

func (r *siteUserRepository) CheckSiteUserExistsBySiteIdAndUserId(siteId, userId int) (bool, error) {
	var siteUser models.SiteUser

	err := r.db.Where("site_id = ? AND user_id = ?", siteId, userId).First(&siteUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
