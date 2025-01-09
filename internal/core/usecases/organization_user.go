package usecases

import (
	"server/internal/core/models"
	"server/internal/core/repositories"
)

type OrganizationUserUsecase interface {
	CreateOrganizationUser(organizationUser *models.OrganizationUser) (*models.OrganizationUser, error)
}

type organizationUserUsecase struct {
	organizationRepo     repositories.OrganizationRepository
	organizationUserRepo repositories.OrganizationUserRepository
	organizationLogRepo  repositories.OrganizationLogRepository
}

func NewOrganizationUserUsecase(organizationRepo repositories.OrganizationRepository, organizationUserRepo repositories.OrganizationUserRepository, organizationLogRepo repositories.OrganizationLogRepository) OrganizationUserUsecase {
	return &organizationUserUsecase{organizationRepo, organizationUserRepo, organizationLogRepo}
}

// TODO: Implement this
func (u *organizationUserUsecase) CreateOrganizationUser(organizationUser *models.OrganizationUser) (*models.OrganizationUser, error) {
	txOrganizationUserRepo, err := u.organizationUserRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txOrganizationUserRepo.Rollback()
		}
	}()

	newOrganizationUser, err := txOrganizationUserRepo.CreateOrganizationUser(organizationUser)
	if err != nil {
		txOrganizationUserRepo.Rollback()
		return nil, err
	}

	if err := txOrganizationUserRepo.Commit(); err != nil {
		return nil, err
	}

	return newOrganizationUser, nil
}
