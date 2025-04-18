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
	GetWorkspaceById(id int) (*models.Workspace, error)
	UpdateSite(site *models.Site) (*models.Site, error)
	DeleteSite(site *models.Site) error
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

func (r *siteRepository) GetWorkspaceById(id int) (*models.Workspace, error) {
	workspace := new(models.Workspace)

	err := r.db.Raw(`
		SELECT 
			s.site_id, 
			st.site_parent_id,
			s.site_type_id, 
			s.name, 
			s.description, 
			s.short_description, 
			s.url, 
			s.image_url, 
			s.created_at, 
			s.created_by, 
			s.updated_at, 
			s.updated_by, 
			s.deleted_at
		FROM site_trees st
		JOIN sites s ON st.site_child_id = s.site_id
		WHERE s.site_id = ?`, id).Scan(workspace).Error

	if err != nil {
		return nil, err
	}

	return workspace, nil
}

func (r *siteRepository) UpdateSite(site *models.Site) (*models.Site, error) {
	if err := r.db.Save(site).Error; err != nil {
		return nil, err
	}

	return site, nil
}

func (r *siteRepository) DeleteSite(site *models.Site) error {
	if err := r.db.Where("site_id = ?", site.SiteId).Delete(site).Error; err != nil {
		return err
	}

	return nil
}
