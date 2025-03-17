package usecases

import (
	"server/internal/core/models"
	"server/internal/core/repositories"
)

type SiteMiniAppUserUsecase interface {
	CreateSiteMiniAppUser(siteMiniAppUser []models.SiteMiniAppUser) ([]models.SiteMiniAppUser, error)
	GetListSiteMiniAppUserBySiteId(siteId int) ([]models.SiteMiniAppUserJoinTable, error)
	DeleteSiteMiniAppUserBySiteIdAndUserId(siteMiniAppUser []models.SiteMiniAppUser) error
}

type siteMiniAppUserUsecase struct {
	siteMiniAppUser repositories.SiteMiniAppUserRepository
	siteRepo        repositories.SiteRepository
	siteLogRepo     repositories.SiteLogRepository
	userRepo        repositories.UserRepository
}

func NewSiteMiniAppUserUsecase(siteMiniAppUser repositories.SiteMiniAppUserRepository, siteRepo repositories.SiteRepository, siteLogRepo repositories.SiteLogRepository, userRepo repositories.UserRepository) SiteMiniAppUserUsecase {
	return &siteMiniAppUserUsecase{
		siteMiniAppUser: siteMiniAppUser,
		siteRepo:        siteRepo,
		siteLogRepo:     siteLogRepo,
		userRepo:        userRepo,
	}
}

func (u *siteMiniAppUserUsecase) CreateSiteMiniAppUser(siteMiniAppUser []models.SiteMiniAppUser) ([]models.SiteMiniAppUser, error) {
	txSiteMiniAppUserRepo, err := u.siteMiniAppUser.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txSiteMiniAppUserRepo.Rollback()
		}
	}()

	var createdSiteMiniAppUser []models.SiteMiniAppUser

	for _, siteMiniAppUser := range siteMiniAppUser {
		exists, err := txSiteMiniAppUserRepo.CheckSiteMiniAppUserExistsBySiteIdAndUserId(siteMiniAppUser.SiteMiniAppId, siteMiniAppUser.UserId)
		if err != nil {
			txSiteMiniAppUserRepo.Rollback()
			return nil, err
		}

		if exists {
			continue
		}

		createdUser, err := txSiteMiniAppUserRepo.CreateSiteMiniAppUser(&siteMiniAppUser)
		if err != nil {
			txSiteMiniAppUserRepo.Rollback()
			return nil, err
		}
		createdSiteMiniAppUser = append(createdSiteMiniAppUser, *createdUser)
	}

	if err := txSiteMiniAppUserRepo.Commit(); err != nil {
		return nil, err
	}

	return createdSiteMiniAppUser, nil
}

func (u *siteMiniAppUserUsecase) GetListSiteMiniAppUserBySiteId(siteId int) ([]models.SiteMiniAppUserJoinTable, error) {
	return u.siteMiniAppUser.GetListSiteMiniAppUserBySiteId(siteId)
}

func (u *siteMiniAppUserUsecase) DeleteSiteMiniAppUserBySiteIdAndUserId(siteMiniAppUser []models.SiteMiniAppUser) error {
	txSiteMiniAppUserRepo, err := u.siteMiniAppUser.BeginLog()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			txSiteMiniAppUserRepo.Rollback()
		}
	}()

	for _, siteMiniAppUser := range siteMiniAppUser {
		err := txSiteMiniAppUserRepo.DeleteSiteMiniAppUserBySiteIdAndUserId(&siteMiniAppUser)
		if err != nil {
			txSiteMiniAppUserRepo.Rollback()
			return err
		}
	}

	if err := txSiteMiniAppUserRepo.Commit(); err != nil {
		return err
	}

	return nil
}
