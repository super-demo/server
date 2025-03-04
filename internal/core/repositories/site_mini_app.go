package repositories

import (
	"server/internal/core/models"

	"gorm.io/gorm"
)

type SiteMiniAppRepository interface {
	BeginLog() (SiteMiniAppRepository, error)
	Commit() error
	Rollback() error
	CreateSiteMiniApp(siteMiniApp *models.SiteMiniApp) (*models.SiteMiniApp, error)
	CheckSiteMiniAppExistsBySlug(slug string) (bool, error)
	GetListSiteMiniAppBySiteId(siteId int) ([]models.SiteMiniApp, error)
	GetSiteMiniAppById(id int) (*models.SiteMiniApp, error)
	GetSiteMiniAppBySlug(slug string) (*models.SiteMiniApp, error)
	UpdateSiteMiniApp(siteMiniApp *models.SiteMiniApp) (*models.SiteMiniApp, error)
	DeleteSiteMiniApp(siteMiniApp *models.SiteMiniApp) error
}

type siteMiniAppRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewSiteMiniAppRepository(db *gorm.DB) SiteMiniAppRepository {
	return &siteMiniAppRepository{db: db}
}

func (r *siteMiniAppRepository) BeginLog() (SiteMiniAppRepository, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &siteMiniAppRepository{db: r.db, tx: tx}, nil
}

func (r *siteMiniAppRepository) Commit() error {
	return r.tx.Commit().Error
}

func (r *siteMiniAppRepository) Rollback() error {
	return r.tx.Rollback().Error
}

func (r *siteMiniAppRepository) CreateSiteMiniApp(siteMiniApp *models.SiteMiniApp) (*models.SiteMiniApp, error) {
	if err := r.db.Create(siteMiniApp).Error; err != nil {
		return nil, err
	}

	return siteMiniApp, nil
}

func (r *siteMiniAppRepository) CheckSiteMiniAppExistsBySlug(slug string) (bool, error) {
	var siteMiniApp models.SiteMiniApp

	err := r.db.Where("slug = ?", slug).First(&siteMiniApp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (r *siteMiniAppRepository) GetListSiteMiniAppBySiteId(siteId int) ([]models.SiteMiniApp, error) {
	var siteMiniApps []models.SiteMiniApp

	err := r.db.Where("site_id = ?", siteId).Find(&siteMiniApps).Error
	if err != nil {
		return nil, err
	}

	return siteMiniApps, nil
}

func (r *siteMiniAppRepository) GetSiteMiniAppById(id int) (*models.SiteMiniApp, error) {
	var siteMiniApp models.SiteMiniApp

	err := r.db.Where("id = ?", id).First(&siteMiniApp).Error
	if err != nil {
		return nil, err
	}

	return &siteMiniApp, nil
}

func (r *siteMiniAppRepository) GetSiteMiniAppBySlug(slug string) (*models.SiteMiniApp, error) {
	var siteMiniApp models.SiteMiniApp

	err := r.db.Where("slug = ?", slug).First(&siteMiniApp).Error
	if err != nil {
		return nil, err
	}

	return &siteMiniApp, nil
}

func (r *siteMiniAppRepository) UpdateSiteMiniApp(siteMiniApp *models.SiteMiniApp) (*models.SiteMiniApp, error) {
	if err := r.db.Save(siteMiniApp).Error; err != nil {
		return nil, err
	}

	return siteMiniApp, nil
}

func (r *siteMiniAppRepository) DeleteSiteMiniApp(siteMiniApp *models.SiteMiniApp) error {
	if err := r.db.Where("site_mini_app_id = ?", siteMiniApp.SiteMiniAppId).Delete(siteMiniApp).Error; err != nil {
		return err
	}

	return nil
}
