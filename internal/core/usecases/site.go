package usecases

import (
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/repositories"
	"time"
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
	CreatePeopleRole(request *models.CreatePeopleRoleRequest, requesterUserId int) (*models.PeopleRole, error)
	GetListPeopleRole(siteId int) ([]models.PeopleRole, error)
	UpdatePeopleRole(role *models.PeopleRole, requesterUserId int) (*models.PeopleRole, error)
	DeletePeopleRole(role *models.PeopleRole, requesterUserId int) error
}

type siteUsecase struct {
	siteRepo       repositories.SiteRepository
	siteTreeRepo   repositories.SiteTreeRepository
	siteUserRepo   repositories.SiteUserRepository
	siteLogRepo    repositories.SiteLogRepository
	peopleRoleRepo repositories.PeopleRoleRepository
}

func NewSiteUsecase(siteRepo repositories.SiteRepository, siteTreeRepo repositories.SiteTreeRepository, siteUserRepo repositories.SiteUserRepository, siteLogRepo repositories.SiteLogRepository, peopleRoleRepo repositories.PeopleRoleRepository) SiteUsecase {
	return &siteUsecase{
		siteRepo:       siteRepo,
		siteTreeRepo:   siteTreeRepo,
		siteUserRepo:   siteUserRepo,
		siteLogRepo:    siteLogRepo,
		peopleRoleRepo: peopleRoleRepo,
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
		SiteId:          newSite.SiteId,
		UserId:          requesterUserId,
		SiteUserLevelId: repositories.SuperAdminUserLevel.UserLevelId,
		IsActive:        true,
		CreatedBy:       requesterUserId,
		UpdatedBy:       requesterUserId,
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

	parentSiteUsers, err := u.siteUserRepo.GetListSiteUserBySiteId(request.SiteParentId)

	for _, parentSiteUser := range parentSiteUsers {

		newSiteUser := &models.SiteUser{
			SiteId:    newSite.SiteId,
			UserId:    parentSiteUser.UserId,
			CreatedBy: requesterUserId,
			UpdatedBy: requesterUserId,
		}

		if parentSiteUser.SiteUserLevelId == 3 {
			newSiteUser.SiteUserLevelId = 4
		} else {
			newSiteUser.SiteUserLevelId = parentSiteUser.SiteUserLevelId
		}

		if _, err := u.siteUserRepo.CreateSiteUser(newSiteUser); err != nil {
			return nil, err
		}
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

func (u *siteUsecase) CreatePeopleRole(request *models.CreatePeopleRoleRequest, requesterUserId int) (*models.PeopleRole, error) {
	txPeopleRoleRepo, err := u.peopleRoleRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txPeopleRoleRepo.Rollback()
		}
	}()

	exists, err := u.peopleRoleRepo.CheckRoleExistsByName(request.Slug)
	if err != nil {
		txPeopleRoleRepo.Rollback()
		return nil, err
	}
	if exists {
		txPeopleRoleRepo.Rollback()
		return nil, app.ErrNameExist
	}

	role := &models.PeopleRole{
		Slug:        request.Slug,
		Description: request.Description,
		SiteId:      request.SiteId,
		CreatedBy:   requesterUserId,
		UpdatedBy:   requesterUserId,
	}

	newRole, err := txPeopleRoleRepo.CreatePeopleRole(role)
	if err != nil {
		txPeopleRoleRepo.Rollback()
		return nil, err
	}

	if err := txPeopleRoleRepo.Commit(); err != nil {
		return nil, err
	}

	siteLog := &models.SiteLog{
		SiteId:    1,
		Action:    "Created",
		Detail:    "Created people role " + role.Slug,
		CreatedBy: requesterUserId,
	}

	if _, err := u.siteLogRepo.CreateSiteLog(siteLog); err != nil {
		return nil, err
	}

	return newRole, nil
}

func (u *siteUsecase) GetListPeopleRole(siteId int) ([]models.PeopleRole, error) {
	return u.peopleRoleRepo.GetRoleListBySiteId(siteId)
}

func (u *siteUsecase) UpdatePeopleRole(role *models.PeopleRole, requesterUserId int) (*models.PeopleRole, error) {
	txPeopleRoleRepo, err := u.peopleRoleRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txPeopleRoleRepo.Rollback()
		}
	}()

	exists, err := txPeopleRoleRepo.CheckRoleExistsByName(role.Slug)
	if err != nil {
		txPeopleRoleRepo.Rollback()
		return nil, err
	}
	if exists {
		txPeopleRoleRepo.Rollback()
		return nil, app.ErrNameExist
	}

	role.CreatedAt = time.Now()
	role.CreatedBy = requesterUserId
	role.UpdatedBy = requesterUserId
	role.UpdatedAt = time.Now()
	newRole, err := txPeopleRoleRepo.UpdateRole(role)
	if err != nil {
		txPeopleRoleRepo.Rollback()
		return nil, err
	}

	if err := txPeopleRoleRepo.Commit(); err != nil {
		return nil, err
	}

	siteLog := &models.SiteLog{
		SiteId:    1,
		Action:    "Updated",
		Detail:    "Updated people role " + role.Slug,
		CreatedBy: requesterUserId,
	}

	if _, err := u.siteLogRepo.CreateSiteLog(siteLog); err != nil {
		return nil, err
	}

	return newRole, nil
}

func (u *siteUsecase) DeletePeopleRole(role *models.PeopleRole, requesterUserId int) error {
	txPeopleRoleRepo, err := u.peopleRoleRepo.BeginLog()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			txPeopleRoleRepo.Rollback()
		}
	}()

	if err := txPeopleRoleRepo.DeleteRole(role); err != nil {
		txPeopleRoleRepo.Rollback()
		return err
	}

	if err := txPeopleRoleRepo.Commit(); err != nil {
		txPeopleRoleRepo.Rollback()
		return err
	}

	return nil
}
