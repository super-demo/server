package usecases

import (
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/repositories"
)

type SiteMiniAppUserUsecase interface {
	CreateSiteMiniAppUserWithoutSign(request []models.CreateSiteMiniAppUserWithoutSignRequest, requesterUserId int) ([]models.SiteMiniAppUser, error)
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

func (u *siteMiniAppUserUsecase) CreateSiteMiniAppUserWithoutSign(request []models.CreateSiteMiniAppUserWithoutSignRequest, requesterUserId int) ([]models.SiteMiniAppUser, error) {
	txUserRepo, err := u.userRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txUserRepo.Rollback()
		}
	}()

	var createdUsers []models.SiteMiniAppUser

	for _, user := range request {
		exists, err := txUserRepo.CheckUserExistsByEmail(user.Email)
		if err != nil {
			txUserRepo.Rollback()
			return nil, err
		}

		var requestUser = &models.User{
			UserLevelId: user.UserLevelId,
			Name:        "",
			Email:       user.Email,
		}

		var newUser *models.User
		if exists {
			newUser, err = txUserRepo.GetUserByEmail(user.Email)
		} else {
			newUser, err = txUserRepo.CreateUser(requestUser)
		}

		if err != nil {
			txUserRepo.Rollback()
			return nil, err
		}

		siteMiniAppUser := &models.SiteMiniAppUser{
			SiteMiniAppId: user.SiteMiniAppId,
			UserId:        newUser.UserId,
			CreatedBy:     requesterUserId,
			UpdatedBy:     requesterUserId,
		}

		exists, err = u.siteMiniAppUser.CheckSiteMiniAppUserExistsBySiteIdAndUserId(siteMiniAppUser.SiteMiniAppId, siteMiniAppUser.UserId)
		if err != nil {
			txUserRepo.Rollback()
			return nil, err
		}
		if exists {
			txUserRepo.Rollback()
			return nil, app.ErrNameExist
		}

		newSiteMiniAppUser, err := u.siteMiniAppUser.CreateSiteMiniAppUser(siteMiniAppUser)
		if err != nil {
			txUserRepo.Rollback()
			return nil, err
		}

		createdUsers = append(createdUsers, *newSiteMiniAppUser)
	}

	if err := txUserRepo.Commit(); err != nil {
		return nil, err
	}

	return createdUsers, nil
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
