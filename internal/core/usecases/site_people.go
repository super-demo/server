package usecases

import (
	"fmt"
	"log"
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/repositories"
)

type SitePeopleUsecase interface {
	CreateSitePeople(users []models.CreateSitePeopleRequest, requesterUserId int) ([]models.SitePeople, error)
	BulkImportUserWithoutSign(siteId int, users []models.BulkImportUser, requesterUserId int) (*models.BulkImportResponse, error)
	GetListSitePeopleBySiteId(siteId int) ([]models.SitePeopleJoinTable, error)
	DeleteSitePeople(sitePeople *models.SitePeople, requesterUserId int) error
}

type sitePeopleUsecase struct {
	siteUserRepo   repositories.SiteUserRepository
	siteRepo       repositories.SiteRepository
	siteLogRepo    repositories.SiteLogRepository
	userRepo       repositories.UserRepository
	sitePeopleRepo repositories.SitePeopleRepository
	siteTreeRepo   repositories.SiteTreeRepository
}

func NewSitePeopleUsecase(siteUserRepo repositories.SiteUserRepository, siteRepo repositories.SiteRepository, siteLogRepo repositories.SiteLogRepository, userRepo repositories.UserRepository, sitePeopleRepo repositories.SitePeopleRepository, siteTreeRepo repositories.SiteTreeRepository) SitePeopleUsecase {
	return &sitePeopleUsecase{
		siteUserRepo:   siteUserRepo,
		siteRepo:       siteRepo,
		siteLogRepo:    siteLogRepo,
		userRepo:       userRepo,
		sitePeopleRepo: sitePeopleRepo,
		siteTreeRepo:   siteTreeRepo,
	}
}

func (u *sitePeopleUsecase) CreateSitePeople(users []models.CreateSitePeopleRequest, requesterUserId int) ([]models.SitePeople, error) {
	txUserRepo, err := u.userRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txUserRepo.Rollback()
		}
	}()

	var createdPeople []models.SitePeople

	for _, user := range users {
		exists, err := txUserRepo.CheckUserExistsByEmail(user.Email)
		if err != nil {
			txUserRepo.Rollback()
			return nil, err
		}

		var requestUser = &models.User{
			SubRoleId: user.SubRoleId,
			SiteId:    1,
			Name:      "",
			Email:     user.Email,
		}

		var newUser *models.User
		if exists {
			log.Println("requestUser exists", requestUser)
			newUser, err = txUserRepo.GetUserByEmail(user.Email)
			if err != nil {
				txUserRepo.Rollback()
				return nil, err
			}

			newUser.SubRoleId = user.SubRoleId
			newUser, err = txUserRepo.UpdateUser(newUser)
			if err != nil {
				txUserRepo.Rollback()
				return nil, err
			}

			log.Println("newUser", newUser)
		} else {
			requestUser.UserLevelId = repositories.PeopleUserLevel.UserLevelId
			log.Println("requestUser", requestUser)
			newUser, err = txUserRepo.CreateUser(requestUser)
			if err != nil {
				txUserRepo.Rollback()
				return nil, err
			}
		}

		sitePeople := &models.SitePeople{
			SiteId:    user.SiteId,
			UserId:    newUser.UserId,
			CreatedBy: requesterUserId,
			UpdatedBy: requesterUserId,
		}

		exists, err = u.sitePeopleRepo.CheckSiteUserExistsBySiteIdAndUserId(sitePeople.SiteId, sitePeople.UserId)
		if err != nil {
			txUserRepo.Rollback()
			return nil, err
		}

		if exists {
			txUserRepo.Rollback()
			return nil, app.ErrNameExist
		}

		parents, err := u.siteTreeRepo.GetSiteParentsBySiteId(user.SiteId)
		if err != nil {
			txUserRepo.Rollback()
			return nil, err
		}

		for _, parent := range parents {
			sitePeople := &models.SitePeople{
				SiteId:    parent.SiteParentId,
				UserId:    newUser.UserId,
				CreatedBy: requesterUserId,
				UpdatedBy: requesterUserId,
			}

			fmt.Println("sitePeople", sitePeople.SiteId, sitePeople.UserId)
			exists, err = u.sitePeopleRepo.CheckSiteUserExistsBySiteIdAndUserId(sitePeople.SiteId, sitePeople.UserId)
			if err != nil {
				txUserRepo.Rollback()
				return nil, err
			}

			if exists {
				continue
			}

			createdSitePeople, err := u.sitePeopleRepo.CreateSitePeople(sitePeople)
			if err != nil {
				txUserRepo.Rollback()
				return nil, err
			}

			createdPeople = append(createdPeople, *createdSitePeople)
		}

		createdSitePeople, err := u.sitePeopleRepo.CreateSitePeople(sitePeople)
		if err != nil {
			txUserRepo.Rollback()
			return nil, err
		}
		createdPeople = append(createdPeople, *createdSitePeople)

		log.Println("createdPeople", createdPeople)
	}

	err = txUserRepo.Commit()
	if err != nil {
		return nil, err
	}

	return createdPeople, nil
}

func (u *sitePeopleUsecase) BulkImportUserWithoutSign(siteId int, users []models.BulkImportUser, requesterUserId int) (*models.BulkImportResponse, error) {
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
				Name:     user.Name,
				Nickname: user.Nickname,
				Email:    user.Email,
				Message:  app.ErrNameExist.Error(),
			})
			continue
		}

		siteUser.CreatedBy = requesterUserId
		siteUser.UpdatedBy = requesterUserId
		_, err = txSiteUserRepo.CreateSiteUser(siteUser)
		if err != nil {
			result.Failures = append(result.Failures, models.BulkImportFailure{
				Name:     user.Name,
				Nickname: user.Nickname,
				Email:    user.Email,
				Message:  err.Error(),
			})
			continue
		}

		result.SuccessCount++
	}

	txSiteUserRepo.Commit()
	return result, nil
}

func (u *sitePeopleUsecase) GetListSitePeopleBySiteId(siteId int) ([]models.SitePeopleJoinTable, error) {
	return u.sitePeopleRepo.GetListSitePeopleBySiteId(siteId)
}

func (u *sitePeopleUsecase) DeleteSitePeople(sitePeople *models.SitePeople, requesterUserId int) error {
	txPeopleRoleRepo, err := u.sitePeopleRepo.BeginLog()
	if err != nil {
		return err
	}

	err = txPeopleRoleRepo.DeleteSiteUserBySiteIdAndUserId(sitePeople)
	if err != nil {
		txPeopleRoleRepo.Rollback()
		return err
	}

	txPeopleRoleRepo.Commit()
	return nil
}
