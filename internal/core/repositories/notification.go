package repositories

import (
	"server/internal/core/models"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	BeginLog() (NotificationRepository, error)
	Commit() error
	Rollback() error
	CreateNotification(notification *models.Notification) (*models.Notification, error)
	GetNotificationById(notificationId int) (*models.Notification, error)
	GetListNotificationByUserId(userId int) (*[]models.Notification, error)
}

type notificationRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db: db}
}

func (r *notificationRepository) BeginLog() (NotificationRepository, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &notificationRepository{db: r.db, tx: tx}, nil
}

func (r *notificationRepository) Commit() error {
	return r.tx.Commit().Error
}

func (r *notificationRepository) Rollback() error {
	return r.tx.Rollback().Error
}

func (r *notificationRepository) CreateNotification(notification *models.Notification) (*models.Notification, error) {
	if err := r.db.Create(notification).Error; err != nil {
		return notification, err
	}

	return notification, nil
}

func (r *notificationRepository) GetNotificationById(notificationId int) (*models.Notification, error) {
	var notification models.Notification
	if err := r.db.Where("notification_id = ?", notificationId).First(&notification).Error; err != nil {
		return nil, err
	}

	return &notification, nil
}

func (r *notificationRepository) GetListNotificationByUserId(userId int) (*[]models.Notification, error) {
	var notifications []models.Notification
	if err := r.db.Table("notifications").
		Select("notifications.*").
		Joins("JOIN notification_users ON notifications.notification_id = notification_users.notification_id").
		Where("notification_users.user_id = ?", userId).
		Find(&notifications).Error; err != nil {
		return nil, err
	}

	return &notifications, nil
}
