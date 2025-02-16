package usecases

import (
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/repositories"
)

type SiteUserUsecase interface {
	CreateSiteUserWithoutSign(request *models.CreateSiteUserWithoutSignRequest, requesterUserId int) (*models.SiteUser, error)
}

type siteUserUsecase struct {
	siteUserRepo repositories.SiteUserRepository
	siteRepo     repositories.SiteRepository
	siteLogRepo  repositories.SiteLogRepository
	userRepo     repositories.UserRepository
}

func NewSiteUserUsecase(siteUserRepo repositories.SiteUserRepository, siteRepo repositories.SiteRepository, siteLogRepo repositories.SiteLogRepository, userRepo repositories.UserRepository) SiteUserUsecase {
	return &siteUserUsecase{
		siteUserRepo: siteUserRepo,
		siteRepo:     siteRepo,
		siteLogRepo:  siteLogRepo,
		userRepo:     userRepo,
	}
}

func (u *siteUserUsecase) CreateSiteUserWithoutSign(request *models.CreateSiteUserWithoutSignRequest, requesterUserId int) (*models.SiteUser, error) {
	txUserRepo, err := u.userRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txUserRepo.Rollback()
		}
	}()

	exists, err := txUserRepo.CheckUserExistsByEmail(request.Email)
	if err != nil {
		txUserRepo.Rollback()
		return nil, err
	}

	var requestUser = &models.User{
		UserLevelId: repositories.SuperAdminUserLevel.UserLevelId,
		Name:        "",
		Email:       request.Email,
	}

	var newUser *models.User
	if exists {
		newUser, err = txUserRepo.GetUserByEmail(request.Email)
		if err != nil {
			txUserRepo.Rollback()
			return nil, err
		}
	} else {
		newUser, err = txUserRepo.CreateUser(requestUser)
		if err != nil {
			txUserRepo.Rollback()
			return nil, err
		}
	}

	siteUser := &models.SiteUser{
		SiteId: request.SiteId,
		UserId: newUser.UserId,
	}

	txSiteUserRepo, err := u.siteUserRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txSiteUserRepo.Rollback()
		}
	}()

	exists, err = txSiteUserRepo.CheckSiteUserExistsBySiteIdAndUserId(siteUser.SiteId, siteUser.UserId)
	if err != nil {
		txSiteUserRepo.Rollback()
		return nil, err
	}

	if exists {
		txSiteUserRepo.Rollback()
		return nil, app.ErrNameExist
	}

	siteUser.CreatedBy = requesterUserId
	siteUser.UpdatedBy = requesterUserId
	newSiteUser, err := txSiteUserRepo.CreateSiteUser(siteUser)
	if err != nil {
		txSiteUserRepo.Rollback()
		return nil, err
	}

	txSiteUserRepo.Commit()
	return newSiteUser, nil
}
