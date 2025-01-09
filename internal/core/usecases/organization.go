package usecases

import (
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/repositories"
)

type OrganizationUsecase interface {
	CreateOrganization(organization *models.Organization, requesterUserId int) (*models.Organization, error)
	GetOrganizationById(organizationId, requesterUserId int) (*models.Organization, error)
	GetOrganizationListByUserId(requesterUserId int) (*[]models.Organization, error)
	DeleteOrganization(organizationId, requesterUserId int) error
}

type organizationUsecase struct {
	organizationRepo     repositories.OrganizationRepository
	organizationUserRepo repositories.OrganizationUserRepository
	organizationLogRepo  repositories.OrganizationLogRepository
}

func NewOrganizationUsecase(organizationRepo repositories.OrganizationRepository, organizationUserRepo repositories.OrganizationUserRepository, organizationLogRepo repositories.OrganizationLogRepository) OrganizationUsecase {
	return &organizationUsecase{organizationRepo, organizationUserRepo, organizationLogRepo}
}

func (u *organizationUsecase) CreateOrganization(organization *models.Organization, requesterUserId int) (*models.Organization, error) {
	txOrganizationRepo, err := u.organizationRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txOrganizationRepo.Rollback()
		}
	}()

	exists, err := txOrganizationRepo.CheckOrganizationExistsByName(organization.Name)
	if err != nil {
		txOrganizationRepo.Rollback()
		return nil, err
	}

	if exists {
		txOrganizationRepo.Rollback()
		return nil, app.ErrOrganizationNameExists
	}

	organization.CreatedBy = requesterUserId
	organization.UpdatedBy = requesterUserId
	newOrganization, err := txOrganizationRepo.CreateOrganization(organization)
	if err != nil {
		txOrganizationRepo.Rollback()
		return nil, err
	}

	if err := txOrganizationRepo.Commit(); err != nil {
		return nil, err
	}

	organizationLog := &models.OrganizationLog{
		OrganizationId: newOrganization.OrganizationId,
		Action:         "Create",
		Description:    "Organization Created",
		CreatedBy:      requesterUserId,
	}

	if _, err := u.organizationLogRepo.CreateOrganizationLog(organizationLog); err != nil {
		return nil, err
	}

	organizationUser := &models.OrganizationUser{
		OrganizationId: newOrganization.OrganizationId,
		UserId:         requesterUserId,
		UserLevelId:    repositories.OwnerUserLevel.UserLevelId,
		IsActive:       true,
		CreatedBy:      requesterUserId,
		UpdatedBy:      requesterUserId,
	}

	if _, err := u.organizationUserRepo.CreateOrganizationUser(organizationUser); err != nil {
		return nil, err
	}

	return newOrganization, nil
}

func (u *organizationUsecase) GetOrganizationById(organizationId, requesterUserId int) (*models.Organization, error) {
	organizationUser, err := u.organizationUserRepo.GetOrganizationUserById(requesterUserId)
	if err != nil {
		if organizationUser == nil {
			return nil, app.ErrUnauthorized
		}
	}

	organization, err := u.organizationRepo.GetOrganizationById(organizationId)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (u *organizationUsecase) GetOrganizationListByUserId(requesterUserId int) (*[]models.Organization, error) {
	var organization *[]models.Organization

	organization, err := u.organizationRepo.GetOrganizationListByUserId(requesterUserId)
	if err != nil {
		if organization == nil {
			return organization, nil
		}
		return nil, err
	}

	return organization, nil
}

func (u *organizationUsecase) DeleteOrganization(organizationId, requesterUserId int) error {
	organization, err := u.organizationRepo.GetOrganizationById(organizationId)
	if err != nil {
		return err
	}

	if organization.CreatedBy != requesterUserId {
		return app.ErrUnauthorized
	}

	if err := u.organizationUserRepo.DeleteOrganizationUserByOrganizationId(organizationId); err != nil {
		return err
	}

	if err := u.organizationRepo.DeleteOrganization(organizationId); err != nil {
		return err
	}

	return nil
}
