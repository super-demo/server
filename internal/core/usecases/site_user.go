package usecases

import (
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/repositories"
)

type SiteUserUsecase interface {
	CreateSiteUserWithoutSign(requests []models.CreateSiteUserWithoutSignRequest, requesterUserId int) ([]models.SiteUser, error)
	BulkImportUserWithoutSign(siteId int, users []models.BulkImportUser, requesterUserId int) (*models.BulkImportResponse, error)
	GetListSiteUserBySiteId(siteId int) ([]models.SiteUserJoinTable, error)
	UpdateSiteUser(siteUser *models.SiteUser) (*models.SiteUser, error)
	DeleteSiteUserBySiteIdAndUserId(siteUser *models.SiteUser, requesterUserId int) error
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

func (u *siteUserUsecase) CreateSiteUserWithoutSign(requests []models.CreateSiteUserWithoutSignRequest, requesterUserId int) ([]models.SiteUser, error) {
	txUserRepo, err := u.userRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txUserRepo.Rollback()
		}
	}()

	var createdUsers []models.SiteUser

	for _, user := range requests {
		exists, err := txUserRepo.CheckUserExistsByEmail(user.Email)
		if err != nil {
			txUserRepo.Rollback()
			return nil, err
		}

		var newUser = &models.User{
			SiteId:      user.SiteId,
			UserLevelId: user.SiteUserLevelId,
			Name:        "",
			Email:       user.Email,
		}

		var requestUser *models.User
		if exists {
			requestUser, err = txUserRepo.GetUserByEmail(user.Email)
			if err != nil {
				txUserRepo.Rollback()
				return nil, err
			}
		} else {
			if newUser.SiteId != 1 {
				newUser.UserLevelId = repositories.ViewerUserLevel.UserLevelId
				newUser.SiteId = 1
			}

			requestUser, err = txUserRepo.CreateUser(newUser)
			if err != nil {
				txUserRepo.Rollback()
				return nil, err
			}
		}

		siteUser := &models.SiteUser{
			SiteId:    user.SiteId,
			UserId:    requestUser.UserId,
			CreatedBy: requesterUserId,
			UpdatedBy: requesterUserId,
		}

		if user.SiteId != 1 && user.SiteUserLevelId == 3 {
			siteUser.SiteUserLevelId = repositories.AdminUserLevel.UserLevelId
		} else {
			siteUser.SiteUserLevelId = user.SiteUserLevelId
		}

		exists, err = u.siteUserRepo.CheckSiteUserExistsBySiteIdAndUserId(siteUser.SiteId, siteUser.UserId)
		if err != nil {
			txUserRepo.Rollback()
			return nil, err
		}

		if exists {
			txUserRepo.Rollback()
			return nil, app.ErrNameExist
		}

		createdSiteUser, err := u.siteUserRepo.CreateSiteUser(siteUser)
		if err != nil {
			txUserRepo.Rollback()
			return nil, err
		}

		createdUsers = append(createdUsers, *createdSiteUser)
	}

	err = txUserRepo.Commit()
	if err != nil {
		return nil, err
	}

	return createdUsers, nil
}

func (u *siteUserUsecase) BulkImportUserWithoutSign(siteId int, users []models.BulkImportUser, requesterUserId int) (*models.BulkImportResponse, error) {
	txUserRepo, err := u.userRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txUserRepo.Rollback()
		}
	}()

	txSiteUserRepo, err := u.siteUserRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txSiteUserRepo.Rollback()
		}
	}()

	var result = &models.BulkImportResponse{}
	for _, user := range users {
		exists, err := txUserRepo.CheckUserExistsByEmail(user.Email)
		if err != nil {
			txUserRepo.Rollback()
			return nil, err
		}

		var requestUser = &models.User{
			UserLevelId: repositories.SuperAdminUserLevel.UserLevelId,
			SiteId:      1,
			Name:        user.Name,
			Nickname:    user.Nickname,
			Email:       user.Email,
		}

		var newUser *models.User
		if exists {
			newUser, err = txUserRepo.GetUserByEmail(user.Email)
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
			SiteId: siteId,
			UserId: newUser.UserId,
		}

		exists, err = txSiteUserRepo.CheckSiteUserExistsBySiteIdAndUserId(siteUser.SiteId, siteUser.UserId)
		if err != nil {
			txSiteUserRepo.Rollback()
			return nil, err
		}

		if exists {
			result.Failures = append(result.Failures, models.BulkImportFailure{
				Name:            user.Name,
				Nickname:        user.Nickname,
				Email:           user.Email,
				SiteUserLevelId: user.SiteUserLevelId,
				Message:         app.ErrNameExist.Error(),
			})
			continue
		}

		siteUser.CreatedBy = requesterUserId
		siteUser.UpdatedBy = requesterUserId
		_, err = txSiteUserRepo.CreateSiteUser(siteUser)
		if err != nil {
			result.Failures = append(result.Failures, models.BulkImportFailure{
				Name:            user.Name,
				Nickname:        user.Nickname,
				Email:           user.Email,
				SiteUserLevelId: user.SiteUserLevelId,
				Message:         err.Error(),
			})
			continue
		}

		result.SuccessCount++
	}

	txSiteUserRepo.Commit()
	return result, nil
}

func (u *siteUserUsecase) GetListSiteUserBySiteId(siteId int) ([]models.SiteUserJoinTable, error) {
	return u.siteUserRepo.GetListSiteUserBySiteId(siteId)
}

func (u *siteUserUsecase) UpdateSiteUser(siteUser *models.SiteUser) (*models.SiteUser, error) {
	txSiteUserRepo, err := u.siteUserRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txSiteUserRepo.Rollback()
		}
	}()

	exists, err := txSiteUserRepo.CheckSiteUserExistsBySiteIdAndUserId(siteUser.SiteId, siteUser.UserId)
	if err != nil {
		txSiteUserRepo.Rollback()
		return nil, err
	}

	if !exists {
		txSiteUserRepo.Rollback()
		return nil, app.ErrNotFound
	}

	siteUser.UpdatedBy = siteUser.UserId
	newSiteUser, err := txSiteUserRepo.UpdateSiteUser(siteUser)
	if err != nil {
		txSiteUserRepo.Rollback()
		return nil, err
	}

	if err := txSiteUserRepo.Commit(); err != nil {
		return nil, err
	}

	return newSiteUser, nil
}

func (u *siteUserUsecase) DeleteSiteUserBySiteIdAndUserId(siteUser *models.SiteUser, requesterUserId int) error {
	txSiteUserRepo, err := u.siteUserRepo.BeginLog()
	if err != nil {
		return err
	}

	err = txSiteUserRepo.DeleteSiteUserBySiteIdAndUserId(siteUser)
	if err != nil {
		txSiteUserRepo.Rollback()
		return err
	}

	txSiteUserRepo.Commit()
	return nil
}
