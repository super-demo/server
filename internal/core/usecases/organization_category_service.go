package usecases

import (
	"server/infrastructure/app"
	"server/internal/core/models"
	"server/internal/core/repositories"
	"strconv"
)

type OrganizationCategoryServiceUsecase interface {
	CreateOrganizationCategoryService(organizationCategoryService *models.OrganizationCategoryService, requesterUserId int) (*models.OrganizationCategoryService, error)
}

type organizationCategoryServiceUsecase struct {
	organizationCategoryRepo        repositories.OrganizationCategoryRepository
	organizationCategoryServiceRepo repositories.OrganizationCategoryServiceRepository
	organizationLogRepo             repositories.OrganizationLogRepository
}

func NewOrganizationCategoryServiceUsecase(organizationCategoryRepo repositories.OrganizationCategoryRepository, organizationCategoryServiceRepo repositories.OrganizationCategoryServiceRepository, organizationLogRepo repositories.OrganizationLogRepository) OrganizationCategoryServiceUsecase {
	return &organizationCategoryServiceUsecase{organizationCategoryRepo, organizationCategoryServiceRepo, organizationLogRepo}
}

func (u *organizationCategoryServiceUsecase) CreateOrganizationCategoryService(organizationCategoryService *models.OrganizationCategoryService, requesterUserId int) (*models.OrganizationCategoryService, error) {
	txOrganizationCategoryServiceRepo, err := u.organizationCategoryServiceRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txOrganizationCategoryServiceRepo.Rollback()
		}
	}()

	exists, err := txOrganizationCategoryServiceRepo.CheckOrganizationCategoryServiceExistsById(organizationCategoryService.OrganizationServiceId)
	if err != nil {
		txOrganizationCategoryServiceRepo.Rollback()
		return nil, err
	}

	if exists {
		txOrganizationCategoryServiceRepo.Rollback()
		return nil, app.ErrOrganizationCategoryServiceExists
	}

	organizationCategoryService.CreatedBy = requesterUserId
	organizationCategoryService.UpdatedBy = requesterUserId
	newOrganizationCategoryService, err := txOrganizationCategoryServiceRepo.CreateOrganizationCategoryService(organizationCategoryService)
	if err != nil {
		txOrganizationCategoryServiceRepo.Rollback()
		return nil, err
	}

	if err := txOrganizationCategoryServiceRepo.Commit(); err != nil {
		return nil, err
	}

	organizationLog := &models.OrganizationLog{
		OrganizationId: newOrganizationCategoryService.OrganizationId,
		Action:         "Added",
		Description:    "Added service " + strconv.Itoa(newOrganizationCategoryService.OrganizationServiceId),
		CreatedBy:      requesterUserId,
	}

	if _, err := u.organizationLogRepo.CreateOrganizationLog(organizationLog); err != nil {
		return nil, err
	}

	return newOrganizationCategoryService, nil
}
