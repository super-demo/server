package repositories

import (
	"server/internal/core/models"

	"gorm.io/gorm"
)

type SitePeopleRepository interface {
	BeginLog() (SitePeopleRepository, error)
	Commit() error
	Rollback() error
	CreateSitePeople(sitePeople *models.SitePeople) (*models.SitePeople, error)
	CheckSiteUserExistsBySiteIdAndUserId(siteId, userId int) (bool, error)
	GetListUserBySiteId(siteId int) ([]models.SiteUser, error)
	GetListSitePeopleBySiteId(siteId int) ([]models.SitePeopleJoinTable, error)
	DeleteSiteUserBySiteIdAndUserId(siteUser *models.SiteUser) error
}

type sitePeopleRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewSitePeopleRepository(db *gorm.DB) SitePeopleRepository {
	return &sitePeopleRepository{db: db}
}

func (r *sitePeopleRepository) BeginLog() (SitePeopleRepository, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &sitePeopleRepository{db: r.db, tx: tx}, nil
}

func (r *sitePeopleRepository) Commit() error {
	return r.tx.Commit().Error
}

func (r *sitePeopleRepository) Rollback() error {
	return r.tx.Rollback().Error
}

func (r *sitePeopleRepository) CreateSitePeople(sitePeople *models.SitePeople) (*models.SitePeople, error) {
	if err := r.db.Create(sitePeople).Error; err != nil {
		return nil, err
	}

	return sitePeople, nil
}

func (r *sitePeopleRepository) CheckSiteUserExistsBySiteIdAndUserId(siteId, userId int) (bool, error) {
	var sitePeople models.SitePeople

	err := r.db.Where("site_id = ? AND user_id = ?", siteId, userId).First(&sitePeople).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *sitePeopleRepository) GetListUserBySiteId(siteId int) ([]models.SiteUser, error) {
	var siteUsers []models.SiteUser
	err := r.db.Where("site_id = ?", siteId).Find(&siteUsers).Error
	if err != nil {
		return nil, err
	}
	return siteUsers, nil
}

func (r *sitePeopleRepository) GetListSitePeopleBySiteId(siteId int) ([]models.SitePeopleJoinTable, error) {
	var sitePeople []models.SitePeopleJoinTable
	err := r.db.Table("site_peoples").
		Preload("User").
		Where("site_peoples.site_id = ?", siteId).
		Find(&sitePeople).Error
	if err != nil {
		return nil, err
	}
	return sitePeople, nil
}

func (r *sitePeopleRepository) DeleteSiteUserBySiteIdAndUserId(siteUser *models.SiteUser) error {
	err := r.db.Where("site_id = ? AND user_id = ?", siteUser.SiteId, siteUser.UserId).Delete(&models.SiteUser{}).Error
	if err != nil {
		return err
	}
	return nil
}
