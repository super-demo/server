package repositories

import (
	"server/internal/core/models"

	"gorm.io/gorm"
)

type AnnouncementRepository interface {
	BeginLog() (AnnouncementRepository, error)
	Commit() error
	Rollback() error
	CreateAnnouncement(announcement *models.Announcement) (*models.Announcement, error)
	GetAnnouncementById(announcementId int) (*models.Announcement, error)
	GetListAnnouncementBySiteId(siteId int) ([]models.Announcement, error)
	UpdateAnnouncement(announcement *models.Announcement) (*models.Announcement, error)
	DeleteAnnouncement(announcementId int) error
}

type announcementRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewAnnouncementRepository(db *gorm.DB) AnnouncementRepository {
	return &announcementRepository{db: db}
}

func (r *announcementRepository) BeginLog() (AnnouncementRepository, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &announcementRepository{db: r.db, tx: tx}, nil
}

func (r *announcementRepository) Commit() error {
	return r.tx.Commit().Error
}

func (r *announcementRepository) Rollback() error {
	return r.tx.Rollback().Error
}

func (r *announcementRepository) CreateAnnouncement(announcement *models.Announcement) (*models.Announcement, error) {
	if err := r.db.Create(announcement).Error; err != nil {
		return nil, err
	}

	return announcement, nil
}

func (r *announcementRepository) GetAnnouncementById(announcementId int) (*models.Announcement, error) {
	var announcement models.Announcement
	if err := r.db.Where("announcement_id = ?", announcementId).First(&announcement).Error; err != nil {
		return nil, err
	}

	return &announcement, nil
}

func (r *announcementRepository) GetListAnnouncementBySiteId(siteId int) ([]models.Announcement, error) {
	var announcements []models.Announcement
	if err := r.db.Where("site_id = ?", siteId).Find(&announcements).Error; err != nil {
		return nil, err
	}

	return announcements, nil
}

func (r *announcementRepository) UpdateAnnouncement(announcement *models.Announcement) (*models.Announcement, error) {
	if err := r.db.Save(announcement).Error; err != nil {
		return nil, err
	}

	return announcement, nil
}

func (r *announcementRepository) DeleteAnnouncement(announcementId int) error {
	if err := r.db.Where("announcement_id = ?", announcementId).Delete(&models.Announcement{}).Error; err != nil {
		return err
	}

	return nil
}
