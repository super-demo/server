package usecases

import (
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/repositories"
)

type SiteUsecase interface {
	CreateSite(site *models.Site, requesterUserId int) (*models.Site, error)
	GetListSite() ([]models.Site, error)
	GetListSiteBySiteTypeId(siteTypeId int) ([]models.Site, error)
	GetListSiteWithoutBySiteTypeId(siteTypeId int) ([]models.Site, error)
	GetSiteById(siteId int) (*models.Site, error)
	GetWorkspaceById(siteId int) (*models.Workspace, error)
	CreateSiteWorkspace(request *models.CreateSiteWorkspaceRequest, requesterUserId int) (*models.Site, error)
	UpdateSiteWorkspace(site *models.Site, requesterUserId int) (*models.Site, error)
	DeleteSiteWorkspace(site *models.Site, requesterUserId int) error
}

type siteUsecase struct {
	siteRepo     repositories.SiteRepository
	siteTreeRepo repositories.SiteTreeRepository
	siteUserRepo repositories.SiteUserRepository
	siteLogRepo  repositories.SiteLogRepository
}

func NewSiteUsecase(siteRepo repositories.SiteRepository, siteTreeRepo repositories.SiteTreeRepository, siteUserRepo repositories.SiteUserRepository, siteLogRepo repositories.SiteLogRepository) SiteUsecase {
	return &siteUsecase{
		siteRepo:     siteRepo,
		siteTreeRepo: siteTreeRepo,
		siteUserRepo: siteUserRepo,
		siteLogRepo:  siteLogRepo,
	}
}

func (u *siteUsecase) CreateSite(site *models.Site, requesterUserId int) (*models.Site, error) {
	txSiteRepo, err := u.siteRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txSiteRepo.Rollback()
		}
	}()

	if site.SiteTypeId == 1 {
		exists, err := txSiteRepo.CheckSiteExistsByName(site.Name)
		if err != nil {
			txSiteRepo.Rollback()
			return nil, err
		}
		if exists {
			txSiteRepo.Rollback()
			return nil, app.ErrNameExist
		}
	}

	site.Url = "https://super-office-cms-ecru.vercel.app/" + site.Name
	site.CreatedBy = requesterUserId
	site.UpdatedBy = requesterUserId
	newSite, err := txSiteRepo.CreateSite(site)
	if err != nil {
		txSiteRepo.Rollback()
		return nil, err
	}

	if err := txSiteRepo.Commit(); err != nil {
		return nil, err
	}

	siteLog := &models.SiteLog{
		SiteId:    newSite.SiteId,
		Action:    "Created",
		Detail:    "Created site " + newSite.Name,
		CreatedBy: requesterUserId,
	}

	if _, err := u.siteLogRepo.CreateSiteLog(siteLog); err != nil {
		return nil, err
	}

	siteUser := &models.SiteUser{
		SiteId:    newSite.SiteId,
		UserId:    requesterUserId,
		IsActive:  true,
		CreatedBy: requesterUserId,
		UpdatedBy: requesterUserId,
	}

	if _, err := u.siteUserRepo.CreateSiteUser(siteUser); err != nil {
		return nil, err
	}

	return newSite, nil
}

func (u *siteUsecase) GetListSite() ([]models.Site, error) {
	return u.siteRepo.GetListSite()
}

func (u *siteUsecase) GetListSiteBySiteTypeId(siteTypeId int) ([]models.Site, error) {
	return u.siteRepo.GetListSiteBySiteTypeId(siteTypeId)
}

func (u *siteUsecase) GetListSiteWithoutBySiteTypeId(siteTypeId int) ([]models.Site, error) {
	return u.siteRepo.GetListSiteWithoutBySiteTypeId(siteTypeId)
}

func (u *siteUsecase) GetSiteById(siteId int) (*models.Site, error) {
	return u.siteRepo.GetSiteById(siteId)
}

func (u *siteUsecase) GetWorkspaceById(siteId int) (*models.Workspace, error) {
	return u.siteRepo.GetWorkspaceById(siteId)
}

func (u *siteUsecase) CreateSiteWorkspace(request *models.CreateSiteWorkspaceRequest, requesterUserId int) (*models.Site, error) {
	txSiteRepo, err := u.siteRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txSiteRepo.Rollback()
		}
	}()

	request.Site.SiteTypeId = 1
	request.Site.CreatedBy = requesterUserId
	request.Site.UpdatedBy = requesterUserId
	newSite, err := txSiteRepo.CreateSite(&request.Site)
	if err != nil {
		txSiteRepo.Rollback()
		return nil, err
	}

	if err := txSiteRepo.Commit(); err != nil {
		return nil, err
	}

	if request.SiteParentId != 0 {
		siteTree := &models.SiteTree{
			SiteParentId: request.SiteParentId,
			SiteChildId:  newSite.SiteId,
			CreatedBy:    requesterUserId,
			UpdatedBy:    requesterUserId,
		}

		if _, err := u.siteTreeRepo.CreateSiteTree(siteTree); err != nil {
			return nil, err
		}
	}

	siteLog := &models.SiteLog{
		SiteId:    newSite.SiteId,
		Action:    "Created",
		Detail:    "Created site " + newSite.Name,
		CreatedBy: requesterUserId,
	}

	if _, err := u.siteLogRepo.CreateSiteLog(siteLog); err != nil {
		return nil, err
	}

	siteUser := &models.SiteUser{
		SiteId:    newSite.SiteId,
		UserId:    requesterUserId,
		IsActive:  true,
		CreatedBy: requesterUserId,
		UpdatedBy: requesterUserId,
	}

	if _, err := u.siteUserRepo.CreateSiteUser(siteUser); err != nil {
		return nil, err
	}

	return newSite, nil
}

func (u *siteUsecase) UpdateSiteWorkspace(site *models.Site, requesterUserId int) (*models.Site, error) {
	txSiteRepo, err := u.siteRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txSiteRepo.Rollback()
		}
	}()

	exists, err := txSiteRepo.CheckSiteExistsByName(site.Name)
	if err != nil {
		txSiteRepo.Rollback()
		return nil, err
	}
	if exists {
		txSiteRepo.Rollback()
		return nil, app.ErrNameExist
	}

	site.UpdatedBy = requesterUserId
	newSite, err := txSiteRepo.UpdateSite(site)
	if err != nil {
		txSiteRepo.Rollback()
		return nil, err
	}

	if err := txSiteRepo.Commit(); err != nil {
		return nil, err
	}

	siteLog := &models.SiteLog{
		SiteId:    newSite.SiteId,
		Action:    "Updated",
		Detail:    "Updated site " + newSite.Name,
		CreatedBy: requesterUserId,
	}

	if _, err := u.siteLogRepo.CreateSiteLog(siteLog); err != nil {
		return nil, err
	}

	return newSite, nil
}

func (u *siteUsecase) DeleteSiteWorkspace(site *models.Site, requesterUserId int) error {
	txSiteRepo, err := u.siteRepo.BeginLog()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			txSiteRepo.Rollback()
		}
	}()

	if err := txSiteRepo.DeleteSite(site); err != nil {
		txSiteRepo.Rollback()
		return err
	}

	if err := txSiteRepo.Commit(); err != nil {
		txSiteRepo.Rollback()
		return err
	}

	return nil
}
