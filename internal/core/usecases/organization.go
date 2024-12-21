package usecases

import (
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/repositories"
)

type OrganizationUsecase interface {
	CreateOrganization(organization *models.Organization, requesterUserId int) (*models.Organization, error)
	GetOrganizationListByUserId(requesterUserId int) (*[]models.Organization, error)
}

type organizationUsecase struct {
	organizationRepo     repositories.OrganizationRepository
	organizationUserRepo repositories.OrganizationUserRepository
}

func NewOrganizationUsecase(organizationRepo repositories.OrganizationRepository, organizationUserRepo repositories.OrganizationUserRepository) OrganizationUsecase {
	return &organizationUsecase{organizationRepo, organizationUserRepo}
}

func (u *organizationUsecase) CreateOrganization(organization *models.Organization, requesterUserId int) (*models.Organization, error) {
	organization.CreatedBy = requesterUserId
	organization.UpdatedBy = requesterUserId
	organization, err := u.organizationRepo.CreateOrganization(organization)
	if err != nil {
		return nil, err
	}

	organizationUser := &models.OrganizationUser{
		OrganizationId: organization.OrganizationId,
		UserId:         requesterUserId,
		UserLevelId:    1,
		IsActive:       true,
		CreatedBy:      requesterUserId,
		UpdatedBy:      requesterUserId,
	}

	if _, err := u.organizationUserRepo.CreateOrganizationUser(organizationUser); err != nil {
		return nil, err
	}

	return organization, nil
}

func (u *organizationUsecase) GetOrganizationListByUserId(requesterUserId int) (*[]models.Organization, error) {
	var organization *[]models.Organization

	organizationUser, err := u.organizationUserRepo.GetOrganizationUserById(requesterUserId)
	if err != nil {
		if organizationUser == nil {
			return nil, app.ErrUnauthorized
		}
	}

	organization, err = u.organizationRepo.GetOrganizationListByUserId(requesterUserId)
	if err != nil {
		if organization == nil {
			return organization, nil
		}
		return nil, err
	}

	return organization, nil
}
