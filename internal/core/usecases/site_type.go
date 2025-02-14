package usecases

import (
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/repositories"
)

type SiteTypeUsecase interface {
	CreateSiteType(siteType *models.SiteType, requesterUserId int) (*models.SiteType, error)
	GetListSiteType() ([]models.SiteType, error)
}

type siteTypeUsecase struct {
	siteTypeRepo repositories.SiteTypeRepository
	siteRepo     repositories.SiteRepository
	siteUserRepo repositories.SiteUserRepository
	siteLogRepo  repositories.SiteLogRepository
}

func NewSiteTypeUsecase(siteTypeRepo repositories.SiteTypeRepository, siteRepo repositories.SiteRepository, siteUserRepo repositories.SiteUserRepository, siteLogRepo repositories.SiteLogRepository) SiteTypeUsecase {
	return &siteTypeUsecase{
		siteTypeRepo: siteTypeRepo,
		siteRepo:     siteRepo,
		siteUserRepo: siteUserRepo,
		siteLogRepo:  siteLogRepo,
	}
}

func (u *siteTypeUsecase) CreateSiteType(siteType *models.SiteType, requesterUserId int) (*models.SiteType, error) {
	txSiteTypeRepo, err := u.siteTypeRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txSiteTypeRepo.Rollback()
		}
	}()

	exists, err := txSiteTypeRepo.CheckSiteTypeExistsByName(siteType.Slug)
	if err != nil {
		txSiteTypeRepo.Rollback()
		return nil, err
	}

	if exists {
		txSiteTypeRepo.Rollback()
		return nil, app.ErrNameExist
	}

	newSiteType, err := txSiteTypeRepo.CreateSiteType(siteType)
	if err != nil {
		txSiteTypeRepo.Rollback()
		return nil, err
	}

	if err := txSiteTypeRepo.Commit(); err != nil {
		return nil, err
	}

	return newSiteType, nil
}

func (u *siteTypeUsecase) GetListSiteType() ([]models.SiteType, error) {
	return u.siteTypeRepo.GetListSiteType()
}
