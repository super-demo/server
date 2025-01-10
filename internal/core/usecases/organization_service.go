package usecases

import (
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/repositories"
)

type OrganizationServiceUsecase interface {
	CreateOrganizationService(organizationService *models.OrganizationService, requesterUserId int) (*models.OrganizationService, error)
}

type organizationServiceUsecase struct {
	organizationRepo        repositories.OrganizationRepository
	organizationServiceRepo repositories.OrganizationServiceRepository
	organizationLogRepo     repositories.OrganizationLogRepository
}

func NewOrganizationServiceUsecase(organizationRepo repositories.OrganizationRepository, organizationServiceRepo repositories.OrganizationServiceRepository, organizationLogRepo repositories.OrganizationLogRepository) OrganizationServiceUsecase {
	return &organizationServiceUsecase{organizationRepo, organizationServiceRepo, organizationLogRepo}
}

func (u *organizationServiceUsecase) CreateOrganizationService(organizationService *models.OrganizationService, requesterUserId int) (*models.OrganizationService, error) {
	txOrganizationServiceRepo, err := u.organizationServiceRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txOrganizationServiceRepo.Rollback()
		}
	}()

	exists, err := txOrganizationServiceRepo.CheckOrganizationServiceExistsByName(organizationService.Slug)
	if err != nil {
		txOrganizationServiceRepo.Rollback()
		return nil, err
	}

	if exists {
		txOrganizationServiceRepo.Rollback()
		return nil, app.ErrOrganizationServiceNameExists
	}

	// TODO: This is a hardcoded value. It should be dynamic. please fix it later.
	organizationService.ServiceId = "9ab8e702-679a-4c06-8f37-92cd11b99695"
	organizationService.CreatedBy = requesterUserId
	organizationService.UpdatedBy = requesterUserId
	newOrganizationService, err := txOrganizationServiceRepo.CreateOrganizationService(organizationService)
	if err != nil {
		txOrganizationServiceRepo.Rollback()
		return nil, err
	}

	if err := txOrganizationServiceRepo.Commit(); err != nil {
		return nil, err
	}

	organizationLog := &models.OrganizationLog{
		OrganizationId: newOrganizationService.OrganizationId,
		Action:         "Add Service",
		Description:    "Service Added to Organization",
		CreatedBy:      requesterUserId,
	}

	if _, err := u.organizationLogRepo.CreateOrganizationLog(organizationLog); err != nil {
		return nil, err
	}

	return newOrganizationService, nil
}
