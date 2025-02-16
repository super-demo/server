package repositories

import (
	"server/infrastructure/app"
	"server/internal/core/models"

	"gorm.io/gorm"
)

type SiteRepository interface {
	BeginLog() (SiteRepository, error)
	Commit() error
	Rollback() error
	CreateSite(site *models.Site) (*models.Site, error)
	CheckSiteExistsByName(name string) (bool, error)
	GetListSite() ([]models.Site, error)
	GetListSiteBySiteTypeId(siteTypeId int) ([]models.Site, error)
	GetListSiteWithoutBySiteTypeId(siteTypeId int) ([]models.Site, error)
	GetSiteById(id int) (*models.Site, error)
	GetSiteByName(name string) (*models.Site, error)
}

type siteRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewSiteRepository(db *gorm.DB) SiteRepository {
	return &siteRepository{db: db}
}

func (r *siteRepository) BeginLog() (SiteRepository, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &siteRepository{db: r.db, tx: tx}, nil
}

func (r *siteRepository) Commit() error {
	return r.tx.Commit().Error
}

func (r *siteRepository) Rollback() error {
	return r.tx.Rollback().Error
}

func (r *siteRepository) CreateSite(site *models.Site) (*models.Site, error) {
	if err := r.db.Create(site).Error; err != nil {
		return nil, err
	}

	return site, nil
}

func (r *siteRepository) CheckSiteExistsByName(name string) (bool, error) {
	var site models.Site

	err := r.db.Where("name = ?", name).First(&site).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, app.ErrNameExist
}

func (r *siteRepository) GetListSite() ([]models.Site, error) {
	var sites []models.Site

	if err := r.db.Find(&sites).Error; err != nil {
		return nil, err
	}

	return sites, nil
}

func (r *siteRepository) GetListSiteBySiteTypeId(siteTypeId int) ([]models.Site, error) {
	var sites []models.Site

	if err := r.db.Where("site_type_id = ?", siteTypeId).Find(&sites).Error; err != nil {
		return nil, err
	}

	return sites, nil
}

func (r *siteRepository) GetListSiteWithoutBySiteTypeId(siteTypeId int) ([]models.Site, error) {
	var sites []models.Site

	if err := r.db.Where("site_type_id != ?", siteTypeId).Find(&sites).Error; err != nil {
		return nil, err
	}

	return sites, nil
}

func (r *siteRepository) GetSiteById(id int) (*models.Site, error) {
	site := new(models.Site)

	if err := r.db.Where("site_id = ?", id).First(site).Error; err != nil {
		return nil, err
	}

	return site, nil
}

func (r *siteRepository) GetSiteByName(name string) (*models.Site, error) {
	site := new(models.Site)

	if err := r.db.Where("name = ?", name).First(site).Error; err != nil {
		return nil, err
	}

	return site, nil
}
