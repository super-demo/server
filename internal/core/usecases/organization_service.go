package usecases

import (
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/repositories"
)

type OrganizationServiceUsecase interface {
	CreateOrganizationService(organizationService *models.OrganizationService, requesterUserId int) (*models.OrganizationService, error)
	DeleteOrganizationService(organizationService *models.OrganizationService, requesterUserId int) error
}

type organizationServiceUsecase struct {
	organizationRepo                repositories.OrganizationRepository
	organizationServiceRepo         repositories.OrganizationServiceRepository
	organizationCategoryServiceRepo repositories.OrganizationCategoryServiceRepository
	organizationLogRepo             repositories.OrganizationLogRepository
}

func NewOrganizationServiceUsecase(organizationRepo repositories.OrganizationRepository, organizationServiceRepo repositories.OrganizationServiceRepository, organizationCategoryServiceRepo repositories.OrganizationCategoryServiceRepository, organizationLogRepo repositories.OrganizationLogRepository) OrganizationServiceUsecase {
	return &organizationServiceUsecase{organizationRepo, organizationServiceRepo, organizationCategoryServiceRepo, organizationLogRepo}
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
		Action:         "Added Service",
		Description:    "Service Added " + newOrganizationService.Slug + " in Organization",
		CreatedBy:      requesterUserId,
	}

	if _, err := u.organizationLogRepo.CreateOrganizationLog(organizationLog); err != nil {
		return nil, err
	}

	return newOrganizationService, nil
}

func (u *organizationServiceUsecase) DeleteOrganizationService(organizationService *models.OrganizationService, requesterUserId int) error {
	txOrganizationServiceRepo, err := u.organizationServiceRepo.BeginLog()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			txOrganizationServiceRepo.Rollback()
		}
	}()

	exists, err := txOrganizationServiceRepo.CheckOrganizationServiceExists(organizationService.OrganizationId, organizationService.OrganizationServiceId)
	if err != nil {
		txOrganizationServiceRepo.Rollback()
		return err
	}

	if !exists {
		txOrganizationServiceRepo.Rollback()
		return app.ErrServiceNotFound
	}

	organizationCategoryService, err := u.organizationCategoryServiceRepo.GetOrganizationCategoryServiceByOrganizationServiceId(organizationService.OrganizationServiceId)
	if err != nil && organizationCategoryService != nil {
		txOrganizationServiceRepo.Rollback()
		return err
	}

	if organizationCategoryService != nil {
		if err := u.organizationCategoryServiceRepo.DeleteOrganizationCategoryService(organizationCategoryService); err != nil {
			txOrganizationServiceRepo.Rollback()
			return err
		}
	}

	if err := txOrganizationServiceRepo.DeleteOrganizationService(organizationService); err != nil {
		txOrganizationServiceRepo.Rollback()
		return err
	}

	if err := txOrganizationServiceRepo.Commit(); err != nil {
		return err
	}

	organizationLog := &models.OrganizationLog{
		OrganizationId: organizationService.OrganizationId,
		Action:         "Deleted",
		Description:    "Deleted service " + organizationService.Slug + " from the organization",
		CreatedBy:      requesterUserId,
	}

	if _, err := u.organizationLogRepo.CreateOrganizationLog(organizationLog); err != nil {
		return err
	}

	return nil
}
