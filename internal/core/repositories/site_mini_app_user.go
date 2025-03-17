package repositories

import (
	"server/internal/core/models"

	"gorm.io/gorm"
)

type SiteMiniAppUserRepository interface {
	BeginLog() (SiteMiniAppUserRepository, error)
	Commit() error
	Rollback() error
	CreateSiteMiniAppUser(siteMiniAppUser *models.SiteMiniAppUser) (*models.SiteMiniAppUser, error)
	CheckSiteMiniAppUserExistsBySiteIdAndUserId(siteId, userId int) (bool, error)
	GetListSiteMiniAppUserBySiteId(siteId int) ([]models.SiteMiniAppUserJoinTable, error)
	DeleteSiteMiniAppUserBySiteIdAndUserId(siteMiniAppUser *models.SiteMiniAppUser) error
}

type siteMiniAppUserRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewSiteMiniAppUserRepository(db *gorm.DB) SiteMiniAppUserRepository {
	return &siteMiniAppUserRepository{db: db}
}

func (r *siteMiniAppUserRepository) BeginLog() (SiteMiniAppUserRepository, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &siteMiniAppUserRepository{db: r.db, tx: tx}, nil
}

func (r *siteMiniAppUserRepository) Commit() error {
	return r.tx.Commit().Error
}

func (r *siteMiniAppUserRepository) Rollback() error {
	return r.tx.Rollback().Error
}

func (r *siteMiniAppUserRepository) CreateSiteMiniAppUser(siteMiniAppUser *models.SiteMiniAppUser) (*models.SiteMiniAppUser, error) {
	if err := r.db.Create(siteMiniAppUser).Error; err != nil {
		return nil, err
	}

	return siteMiniAppUser, nil
}

func (r *siteMiniAppUserRepository) CheckSiteMiniAppUserExistsBySiteIdAndUserId(siteId, userId int) (bool, error) {
	var siteMiniAppUser models.SiteMiniAppUser

	err := r.db.Where("site_mini_app_id = ? AND user_id = ?", siteId, userId).First(&siteMiniAppUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
	}

	return true, err
}

func (r *siteMiniAppUserRepository) GetListSiteMiniAppUserBySiteId(siteId int) ([]models.SiteMiniAppUserJoinTable, error) {
	var siteMiniAppUsers []models.SiteMiniAppUserJoinTable

	err := r.db.Where("site_mini_app_id = ?", siteId).Find(&siteMiniAppUsers).Error
	if err != nil {
		return nil, err
	}

	return siteMiniAppUsers, nil
}

func (r *siteMiniAppUserRepository) DeleteSiteMiniAppUserBySiteIdAndUserId(siteMiniAppUser *models.SiteMiniAppUser) error {
	err := r.db.Where("site_mini_app_id = ? AND user_id = ?", siteMiniAppUser.SiteMiniAppId, siteMiniAppUser.UserId).Delete(&siteMiniAppUser).Error
	if err != nil {
		return err
	}

	return nil
}
