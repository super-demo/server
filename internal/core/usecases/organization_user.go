package usecases

import (
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/repositories"
)

type OrganizationUserUsecase interface {
	CreateOrganizationUser(organizationUser *models.OrganizationUser, requesterUserId int) (*models.OrganizationUser, error)
}

type organizationUserUsecase struct {
	organizationRepo     repositories.OrganizationRepository
	organizationUserRepo repositories.OrganizationUserRepository
	organizationLogRepo  repositories.OrganizationLogRepository
}

func NewOrganizationUserUsecase(organizationRepo repositories.OrganizationRepository, organizationUserRepo repositories.OrganizationUserRepository, organizationLogRepo repositories.OrganizationLogRepository) OrganizationUserUsecase {
	return &organizationUserUsecase{organizationRepo, organizationUserRepo, organizationLogRepo}
}

func (u *organizationUserUsecase) CreateOrganizationUser(organizationUser *models.OrganizationUser, requesterUserId int) (*models.OrganizationUser, error) {
	txOrganizationUserRepo, err := u.organizationUserRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txOrganizationUserRepo.Rollback()
		}
	}()

	exists, err := txOrganizationUserRepo.CheckOrganizationUserExists(organizationUser, requesterUserId)
	if err != nil {
		txOrganizationUserRepo.Rollback()
		return nil, err
	}

	if exists {
		txOrganizationUserRepo.Rollback()
		return nil, app.ErrOrganizationUserExists
	}

	organizationUser.CreatedBy = requesterUserId
	organizationUser.UpdatedBy = requesterUserId
	newOrganizationUser, err := txOrganizationUserRepo.CreateOrganizationUser(organizationUser)
	if err != nil {
		txOrganizationUserRepo.Rollback()
		return nil, err
	}

	if err := txOrganizationUserRepo.Commit(); err != nil {
		return nil, err
	}

	organizationLog := &models.OrganizationLog{
		OrganizationId: newOrganizationUser.OrganizationId,
		Action:         "Invited",
		Description:    "User Invited to Organization",
		CreatedBy:      requesterUserId,
	}

	if _, err := u.organizationLogRepo.CreateOrganizationLog(organizationLog); err != nil {
		return nil, err
	}

	return newOrganizationUser, nil
}
