package usecases

import (
	"server/internal/core/models"
	"server/internal/core/repositories"
)

type AnnouncementUsecase interface {
	CreateAnnouncement(announcement *models.Announcement, requesterUserId int) (*models.Announcement, error)
	GetAnnouncementById(announcementId int) (*models.Announcement, error)
	GetListAnnouncementBySiteId(siteId int) ([]models.Announcement, error)
	UpdateAnnouncement(announcement *models.Announcement) (*models.Announcement, error)
	DeleteAnnouncement(announcementId int) error
}

type announcementUsecase struct {
	announcementRepo     repositories.AnnouncementRepository
	notificationRepo     repositories.NotificationRepository
	notificationUserRepo repositories.NotificationUserRepository
	sitePeopleRepo       repositories.SitePeopleRepository
}

func NewAnnouncementUsecase(announcementRepo repositories.AnnouncementRepository, notificationRepo repositories.NotificationRepository, notificationUserRepo repositories.NotificationUserRepository, sitePeopleRepo repositories.SitePeopleRepository) AnnouncementUsecase {
	return &announcementUsecase{
		announcementRepo:     announcementRepo,
		notificationRepo:     notificationRepo,
		notificationUserRepo: notificationUserRepo,
		sitePeopleRepo:       sitePeopleRepo,
	}
}

func (u *announcementUsecase) CreateAnnouncement(announcement *models.Announcement, requesterUserId int) (*models.Announcement, error) {
	txAnnouncementRepo, err := u.announcementRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txAnnouncementRepo.Rollback()
		}
	}()

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

	txSitePeopleRepo, err := u.sitePeopleRepo.BeginLog()
	if err != nil {
		txAnnouncementRepo.Rollback()
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txSitePeopleRepo.Rollback()
		}
	}()

	notification := &models.Notification{
		SiteId:    announcement.SiteId,
		Action:    announcement.Title,
		Detail:    announcement.ShortDescription,
		IamgeUrl:  announcement.ImageUrl,
		CreatedBy: requesterUserId,
	}

	newNotification, err := txNotificationRepo.CreateNotification(notification)
	if err != nil {
		txAnnouncementRepo.Rollback()
		return nil, err
	}

	users, err := txSitePeopleRepo.GetListUserBySiteId(newNotification.SiteId)
	if err != nil {
		txAnnouncementRepo.Rollback()
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
			txAnnouncementRepo.Rollback()
			return nil, err
		}
	}

	announcement.CreatedBy = requesterUserId
	newAnnouncement, err := txAnnouncementRepo.CreateAnnouncement(announcement)
	if err != nil {
		txAnnouncementRepo.Rollback()
		return nil, err
	}

	if err := txAnnouncementRepo.Commit(); err != nil {
		return nil, err
	}

	if err := txNotificationRepo.Commit(); err != nil {
		return nil, err
	}

	if err := txNotificationUserRepo.Commit(); err != nil {
		return nil, err
	}

	if err := txSitePeopleRepo.Commit(); err != nil {
		return nil, err
	}

	return newAnnouncement, nil
}

func (u *announcementUsecase) GetAnnouncementById(announcementId int) (*models.Announcement, error) {
	announcement, err := u.announcementRepo.GetAnnouncementById(announcementId)
	if err != nil {
		return nil, err
	}

	return announcement, nil
}

func (u *announcementUsecase) GetListAnnouncementBySiteId(siteId int) ([]models.Announcement, error) {
	announcements, err := u.announcementRepo.GetListAnnouncementBySiteId(siteId)
	if err != nil {
		return nil, err
	}

	return announcements, nil
}

func (u *announcementUsecase) UpdateAnnouncement(announcement *models.Announcement) (*models.Announcement, error) {
	txAnnouncementRepo, err := u.announcementRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txAnnouncementRepo.Rollback()
		}
	}()

	updatedAnnouncement, err := txAnnouncementRepo.UpdateAnnouncement(announcement)
	if err != nil {
		txAnnouncementRepo.Rollback()
		return nil, err
	}

	if err := txAnnouncementRepo.Commit(); err != nil {
		return nil, err
	}

	return updatedAnnouncement, nil
}

func (u *announcementUsecase) DeleteAnnouncement(announcementId int) error {
	txAnnouncementRepo, err := u.announcementRepo.BeginLog()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			txAnnouncementRepo.Rollback()
		}
	}()

	if err := txAnnouncementRepo.DeleteAnnouncement(announcementId); err != nil {
		txAnnouncementRepo.Rollback()
		return err
	}

	if err := txAnnouncementRepo.Commit(); err != nil {
		return err
	}

	return nil
}
