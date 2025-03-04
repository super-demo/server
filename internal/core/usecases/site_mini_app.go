package usecases

import (
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/repositories"
)

type SiteMiniAppUsecase interface {
	CreateSiteMiniApp(siteMiniApp *models.SiteMiniApp, requesterUserId int) (*models.SiteMiniApp, error)
	GetListSiteMiniAppBySiteId(siteId int) ([]models.SiteMiniApp, error)
	GetSiteMiniAppById(id int) (*models.SiteMiniApp, error)
	UpdateSiteMiniApp(siteMiniApp *models.SiteMiniApp, requesterUserId int) (*models.SiteMiniApp, error)
	DeleteSiteMiniApp(siteMiniApp *models.SiteMiniApp, requesterUserId int) error
}

type siteMiniAppUsecase struct {
	siteMiniAppRepo repositories.SiteMiniAppRepository
	siteRepo        repositories.SiteRepository
	siteTreeRepo    repositories.SiteTreeRepository
	siteLogRepo     repositories.SiteLogRepository
}

func NewSiteMiniAppUsecase(siteMiniAppRepo repositories.SiteMiniAppRepository, siteRepo repositories.SiteRepository, siteTreeRepo repositories.SiteTreeRepository, siteLogRepo repositories.SiteLogRepository) SiteMiniAppUsecase {
	return &siteMiniAppUsecase{
		siteMiniAppRepo: siteMiniAppRepo,
		siteRepo:        siteRepo,
		siteTreeRepo:    siteTreeRepo,
		siteLogRepo:     siteLogRepo,
	}
}

func (u *siteMiniAppUsecase) CreateSiteMiniApp(siteMiniApp *models.SiteMiniApp, requesterUserId int) (*models.SiteMiniApp, error) {
	txSiteMiniAppRepo, err := u.siteMiniAppRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txSiteMiniAppRepo.Rollback()
		}
	}()

	exists, err := txSiteMiniAppRepo.CheckSiteMiniAppExistsBySlug(siteMiniApp.Slug)
	if err != nil {
		txSiteMiniAppRepo.Rollback()
		return nil, err
	}
	if exists {
		txSiteMiniAppRepo.Rollback()
		return nil, app.ErrNameExist
	}

	siteMiniApp.CreatedBy = requesterUserId
	siteMiniApp.UpdatedBy = requesterUserId
	newSiteMiniApp, err := txSiteMiniAppRepo.CreateSiteMiniApp(siteMiniApp)
	if err != nil {
		txSiteMiniAppRepo.Rollback()
		return nil, err
	}

	if err := txSiteMiniAppRepo.Commit(); err != nil {
		txSiteMiniAppRepo.Rollback()
		return nil, err
	}

	siteLog := &models.SiteLog{
		SiteId:    newSiteMiniApp.SiteId,
		Action:    "Created",
		Detail:    "Created site mini-app" + newSiteMiniApp.Slug,
		CreatedBy: requesterUserId,
	}

	if _, err := u.siteLogRepo.CreateSiteLog(siteLog); err != nil {
		return nil, err
	}

	return newSiteMiniApp, nil

}

func (u *siteMiniAppUsecase) GetListSiteMiniAppBySiteId(siteId int) ([]models.SiteMiniApp, error) {
	return u.siteMiniAppRepo.GetListSiteMiniAppBySiteId(siteId)
}

func (u *siteMiniAppUsecase) GetSiteMiniAppById(id int) (*models.SiteMiniApp, error) {
	return u.siteMiniAppRepo.GetSiteMiniAppById(id)
}

func (u *siteMiniAppUsecase) UpdateSiteMiniApp(siteMiniApp *models.SiteMiniApp, requesterUserId int) (*models.SiteMiniApp, error) {
	txSiteMiniAppRepo, err := u.siteMiniAppRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txSiteMiniAppRepo.Rollback()
		}
	}()

	exists, err := txSiteMiniAppRepo.CheckSiteMiniAppExistsBySlug(siteMiniApp.Slug)
	if err != nil {
		txSiteMiniAppRepo.Rollback()
		return nil, err
	}
	if exists {
		txSiteMiniAppRepo.Rollback()
		return nil, app.ErrNameExist
	}

	siteMiniApp.UpdatedBy = requesterUserId
	newSiteMiniApp, err := txSiteMiniAppRepo.UpdateSiteMiniApp(siteMiniApp)
	if err != nil {
		txSiteMiniAppRepo.Rollback()
		return nil, err
	}

	if err := txSiteMiniAppRepo.Commit(); err != nil {
		txSiteMiniAppRepo.Rollback()
		return nil, err
	}

	siteLog := &models.SiteLog{
		SiteId:    newSiteMiniApp.SiteId,
		Action:    "Updated",
		Detail:    "Updated site mini-app" + newSiteMiniApp.Slug,
		CreatedBy: requesterUserId,
	}

	if _, err := u.siteLogRepo.CreateSiteLog(siteLog); err != nil {
		return nil, err
	}

	return newSiteMiniApp, nil
}

func (u *siteMiniAppUsecase) DeleteSiteMiniApp(siteMiniApp *models.SiteMiniApp, requesterUserId int) error {
	txSiteMiniAppRepo, err := u.siteMiniAppRepo.BeginLog()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			txSiteMiniAppRepo.Rollback()
		}
	}()

	err = txSiteMiniAppRepo.DeleteSiteMiniApp(siteMiniApp)
	if err != nil {
		txSiteMiniAppRepo.Rollback()
		return err
	}

	if err := txSiteMiniAppRepo.Commit(); err != nil {
		txSiteMiniAppRepo.Rollback()
		return err
	}

	siteLog := &models.SiteLog{
		SiteId:    siteMiniApp.SiteId,
		Action:    "Deleted",
		Detail:    "Deleted site mini-app" + siteMiniApp.Slug,
		CreatedBy: requesterUserId,
	}

	if _, err := u.siteLogRepo.CreateSiteLog(siteLog); err != nil {
		return err
	}

	return nil
}
