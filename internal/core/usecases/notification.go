package usecases

import (
	"server/internal/core/models"
	"server/internal/core/repositories"
)

type NotificationUsecase interface {
	CreateNotification(notification *models.Notification) (*models.Notification, error)
	GetNotificationById(notificationId int) (*models.Notification, error)
	GetListNotificationByUserId(userId int) (*[]models.Notification, error)
}

type notificationUsecase struct {
	notificationRepo     repositories.NotificationRepository
	notificationUserRepo repositories.NotificationUserRepository
	siteUserRepo         repositories.SiteUserRepository
}

func NewNotificationUsecase(notificationRepo repositories.NotificationRepository, notificationUserRepo repositories.NotificationUserRepository, siteUserRepo repositories.SiteUserRepository) NotificationUsecase {
	return &notificationUsecase{
		notificationRepo:     notificationRepo,
		notificationUserRepo: notificationUserRepo,
		siteUserRepo:         siteUserRepo,
	}
}

func (u *notificationUsecase) CreateNotification(notification *models.Notification) (*models.Notification, error) {
	txNotificationRepo, err := u.notificationRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txNotificationRepo.Rollback()
		}
	}()

	txNotificationUserRepo, err := u.notificationUserRepo.BeginLog()
	if err != nil {
		txNotificationRepo.Rollback()
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txNotificationUserRepo.Rollback()
		}
	}()

	txSiteUserRepo, err := u.siteUserRepo.BeginLog()
	if err != nil {
		txSiteUserRepo.Rollback()
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txSiteUserRepo.Rollback()
		}
	}()

	newNotification, err := txNotificationRepo.CreateNotification(notification)
	if err != nil {
		txNotificationRepo.Rollback()
		return nil, err
	}

	users, err := txSiteUserRepo.GetListUserBySiteId(newNotification.SiteId)
	if err != nil {
		txSiteUserRepo.Rollback()
		return nil, err
	}

	for _, user := range users {
		notificationUser := &models.NotificationUser{
			NotificationId: newNotification.NotificationId,
			UserId:         user.UserId,
			IsRead:         false,
		}
		_, err := txNotificationUserRepo.CreateNotificationUser(notificationUser)
		if err != nil {
			txNotificationUserRepo.Rollback()
			return nil, err
		}
	}

	if err := txNotificationRepo.Commit(); err != nil {
		return nil, err
	}

	if err := txNotificationUserRepo.Commit(); err != nil {
		return nil, err
	}

	if err := txSiteUserRepo.Commit(); err != nil {
		return nil, err
	}

	return newNotification, nil
}

func (u *notificationUsecase) GetNotificationById(notificationId int) (*models.Notification, error) {
	return u.notificationRepo.GetNotificationById(notificationId)
}

func (u *notificationUsecase) GetListNotificationByUserId(userId int) (*[]models.Notification, error) {
	return u.notificationRepo.GetListNotificationByUserId(userId)
}
