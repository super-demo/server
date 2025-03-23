package usecases

import (
	"server/internal/core/models"
	"server/internal/core/repositories"
)

type NotificationUserUsecase interface {
	UpdateNotificationUser(notificationUser *models.NotificationUser) (*models.NotificationUser, error)
}

type notificationUserUsecase struct {
	notificationUserRepo repositories.NotificationUserRepository
	notificationRepo     repositories.NotificationRepository
}

func NewNotificationUserUsecase(notificationUserRepo repositories.NotificationUserRepository, notificationRepo repositories.NotificationRepository) NotificationUserUsecase {
	return &notificationUserUsecase{
		notificationUserRepo: notificationUserRepo,
		notificationRepo:     notificationRepo,
	}
}

func (u *notificationUserUsecase) UpdateNotificationUser(notificationUser *models.NotificationUser) (*models.NotificationUser, error) {
	txNotificationUserRepo, err := u.notificationUserRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txNotificationUserRepo.Rollback()
		}
	}()

	updatedNotificationUser, err := txNotificationUserRepo.UpdateNotificationUser(notificationUser)
	if err != nil {
		txNotificationUserRepo.Rollback()
		return nil, err
	}

	txNotificationUserRepo.Commit()
	return updatedNotificationUser, nil
}
