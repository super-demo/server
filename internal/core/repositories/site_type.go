package repositories

import (
	"server/internal/core/models"

	"gorm.io/gorm"
)

type SiteTypeRepository interface {
	BeginLog() (SiteTypeRepository, error)
	Commit() error
	Rollback() error
	CreateSiteType(siteType *models.SiteType) (*models.SiteType, error)
	CheckSiteTypeExistsByName(name string) (bool, error)
	GetListSiteType() ([]models.SiteType, error)
}

type siteTypeRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewSiteTypeRepository(db *gorm.DB) SiteTypeRepository {
	return &siteTypeRepository{db: db}
}

func (r *siteTypeRepository) BeginLog() (SiteTypeRepository, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &siteTypeRepository{db: r.db, tx: tx}, nil
}

func (r *siteTypeRepository) Commit() error {
	return r.tx.Commit().Error
}

func (r *siteTypeRepository) Rollback() error {
	return r.tx.Rollback().Error
}

func (r *siteTypeRepository) CreateSiteType(siteType *models.SiteType) (*models.SiteType, error) {
	if err := r.db.Create(siteType).Error; err != nil {
		return nil, err
	}

	return siteType, nil
}

func (r *siteTypeRepository) CheckSiteTypeExistsByName(name string) (bool, error) {
	var siteType models.SiteType

	err := r.db.Where("name = ?", name).First(&siteType).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *siteTypeRepository) GetListSiteType() ([]models.SiteType, error) {
	var siteTypes []models.SiteType

	err := r.db.Find(&siteTypes).Error
	if err != nil {
		return nil, err
	}

	return siteTypes, nil
}
