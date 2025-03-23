package repositories

import (
	"server/internal/core/models"

	"gorm.io/gorm"
)

type NotificationUserRepository interface {
	BeginLog() (NotificationUserRepository, error)
	Commit() error
	Rollback() error
	CreateNotificationUser(notificationUser *models.NotificationUser) (*models.NotificationUser, error)
	UpdateNotificationUser(notificationUser *models.NotificationUser) (*models.NotificationUser, error)
}

type notificationUserRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewNotificationUserRepository(db *gorm.DB) NotificationUserRepository {
	return &notificationUserRepository{db: db}
}

func (r *notificationUserRepository) BeginLog() (NotificationUserRepository, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &notificationUserRepository{db: r.db, tx: tx}, nil
}

func (r *notificationUserRepository) Commit() error {
	return r.tx.Commit().Error
}

func (r *notificationUserRepository) Rollback() error {
	return r.tx.Rollback().Error
}

// func (r *notificationUserRepository) CreateNotificationUser(notificationUser *models.NotificationUser) (*models.NotificationUser, error) {
// 	if err := r.db.Create(notificationUser).Error; err != nil {
// 		return nil, err
// 	}

// 	return notificationUser, nil
// }

func (r *notificationUserRepository) CreateNotificationUser(notificationUser *models.NotificationUser) (*models.NotificationUser, error) {
	// Remove explicit setting of notification_user_id if it's currently being set to 0
	// Let the database auto-increment handle the ID assignment
	notificationUser.NotificationUserId = 0 // Set to 0 or remove this line if the struct initializes to 0 by default

	if err := r.db.Create(notificationUser).Error; err != nil {
		return nil, err
	}
	return notificationUser, nil
}

func (r *notificationUserRepository) UpdateNotificationUser(notificationUser *models.NotificationUser) (*models.NotificationUser, error) {
	if err := r.db.Save(notificationUser).Error; err != nil {
		return nil, err
	}

	return notificationUser, nil
}
