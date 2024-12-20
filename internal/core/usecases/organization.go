package usecases

import (
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/repositories"
)

type OrganizationUsecase interface {
	CreateOrganization(organization *models.Organization, requesterUserId int) (*models.Organization, error)
	GetOrganizationListByUserId(id, requesterUserId int) (*[]models.Organization, error)
}

type organizationUsecase struct {
	organizationRepo repositories.OrganizationRepository
}

func NewOrganizationUsecase(organizationRepo repositories.OrganizationRepository) OrganizationUsecase {
	return &organizationUsecase{organizationRepo}
}

func (u *organizationUsecase) CreateOrganization(organization *models.Organization, requesterUserId int) (*models.Organization, error) {
	organization.CreatedBy = requesterUserId
	organization.UpdatedBy = requesterUserId
	organization, err := u.organizationRepo.CreateOrganization(organization)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (u *organizationUsecase) GetOrganizationListByUserId(id, requesterUserId int) (*[]models.Organization, error) {
	var organization *[]models.Organization

	if requesterUserId != id {
		return nil, app.ErrUnauthorized
	}

	organization, err := u.organizationRepo.GetOrganizationListByUserId(id)
	if err != nil {
		if organization == nil {
			return organization, nil
		}
		return nil, err
	}

	return organization, nil
}
